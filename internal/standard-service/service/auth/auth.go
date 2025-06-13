package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/constants"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"github.com/wisaitas/standard-golang/internal/standard-service/utils"
	"github.com/wisaitas/standard-golang/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(req request.LoginRequest) (resp response.LoginResponse, statusCode int, err error)
	Register(req request.RegisterRequest) (resp response.RegisterResponse, statusCode int, err error)
	Logout(userContext pkg.UserContext) (statusCode int, err error)
	RefreshToken(userContext pkg.UserContext) (resp response.LoginResponse, statusCode int, err error)
}

type authService struct {
	userRepository        repository.UserRepository
	userHistoryRepository repository.UserHistoryRepository
	redis                 pkg.Redis
	bcrypt                pkg.Bcrypt
}

func NewAuthService(
	userRepository repository.UserRepository,
	userHistoryRepository repository.UserHistoryRepository,
	redis pkg.Redis,
	bcrypt pkg.Bcrypt,
) AuthService {
	return &authService{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		redis:                 redis,
		bcrypt:                bcrypt,
	}
}

func (r *authService) Login(req request.LoginRequest) (resp response.LoginResponse, statusCode int, err error) {
	user := entity.User{}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("username = ?", req.Username), nil); err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, http.StatusNotFound, pkg.Error(err)
		}

		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return resp, http.StatusUnauthorized, pkg.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(time.Hour * 1)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	tokenData := map[string]interface{}{
		"user_id": user.ID,
	}

	accessToken, err := utils.GenerateToken(tokenData, accessTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	refreshToken, err := utils.GenerateToken(tokenData, refreshTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	sessionData := pkg.UserContext{
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		BirthDate:    user.BirthDate,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	sessionDataJSON, err := json.Marshal(sessionData)
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), string(sessionDataJSON), accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), string(sessionDataJSON), refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.EntityToResponse(accessToken, refreshToken), http.StatusOK, nil
}

func (r *authService) Register(req request.RegisterRequest) (resp response.RegisterResponse, statusCode int, err error) {
	user := entity.User{}

	hashedPassword, err := r.bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	user.Password = string(hashedPassword)

	tm := pkg.NewTxManager(r.userRepository.GetDB())

	txUserRepository := r.userRepository.WithTxManager(tm)
	txUserHistoryRepository := r.userHistoryRepository.WithTxManager(tm)

	if err := txUserHistoryRepository.Create(&entity.UserHistory{
		Action:       constants.Action.Create,
		OldFirstName: user.FirstName,
		OldLastName:  user.LastName,
		OldBirthDate: user.BirthDate,
		OldPassword:  user.Password,
		OldEmail:     user.Email,
		OldVersion:   user.Version,
	}); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := txUserRepository.Create(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return resp, http.StatusBadRequest, pkg.Error(errors.New("username already exists"))
		}

		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	userHistory := entity.UserHistory{
		Action:       constants.Action.Create,
		OldFirstName: user.FirstName,
		OldLastName:  user.LastName,
		OldBirthDate: user.BirthDate,
		OldPassword:  user.Password,
		OldEmail:     user.Email,
		OldVersion:   user.Version,
	}

	if err := txUserHistoryRepository.Create(&userHistory); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := tm.Commit(); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.EntityToResponse(user), http.StatusCreated, nil
}

func (r *authService) Logout(userContext pkg.UserContext) (statusCode int, err error) {
	if err := r.redis.Del(context.Background(), fmt.Sprintf("access_token:%s", userContext.UserID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Del(context.Background(), fmt.Sprintf("refresh_token:%s", userContext.UserID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	return http.StatusOK, nil
}

func (r *authService) RefreshToken(userContext pkg.UserContext) (resp response.LoginResponse, statusCode int, err error) {
	user := entity.User{}
	if err := r.userRepository.GetBy(&user, pkg.NewCondition("username = ?", userContext.Username), nil); err != nil {
		return resp, http.StatusNotFound, pkg.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(time.Hour * 1)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	tokenData := map[string]interface{}{
		"user_id": user.ID,
	}

	accessToken, err := utils.GenerateToken(tokenData, accessTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	refreshToken, err := utils.GenerateToken(tokenData, refreshTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	sessionData := pkg.UserContext{
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		BirthDate:    user.BirthDate,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	sessionDataJSON, err := json.Marshal(sessionData)
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), string(sessionDataJSON), accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), string(sessionDataJSON), refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.EntityToResponse(accessToken, refreshToken), http.StatusOK, nil
}
