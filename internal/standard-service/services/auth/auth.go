package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wisaitas/standard-golang/internal/standard-service/constants"
	"github.com/wisaitas/standard-golang/internal/standard-service/contexts"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/internal/standard-service/repositories"
	"github.com/wisaitas/standard-golang/internal/standard-service/utils"
	"github.com/wisaitas/standard-golang/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(req requests.LoginRequest) (resp responses.LoginResponse, statusCode int, err error)
	Register(req requests.RegisterRequest) (resp responses.RegisterResponse, statusCode int, err error)
	Logout(userContext contexts.UserContext) (statusCode int, err error)
	RefreshToken(userContext contexts.UserContext) (resp responses.LoginResponse, statusCode int, err error)
}

type authService struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	transactionUtil       pkg.TransactionUtil
	redisUtil             pkg.RedisUtil
	bcryptUtil            pkg.BcryptUtil
}

func NewAuthService(
	userRepository repositories.UserRepository,
	userHistoryRepository repositories.UserHistoryRepository,
	transactionUtil pkg.TransactionUtil,
	redisUtil pkg.RedisUtil,
	bcryptUtil pkg.BcryptUtil,
) AuthService {
	return &authService{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		transactionUtil:       transactionUtil,
		redisUtil:             redisUtil,
		bcryptUtil:            bcryptUtil,
	}
}

func (r *authService) Login(req requests.LoginRequest) (resp responses.LoginResponse, statusCode int, err error) {
	user := models.User{}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("username = ?", req.Username), nil); err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, http.StatusNotFound, pkg.Error(err)
		}

		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.bcryptUtil.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
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

	sessionData := contexts.UserContext{
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

	if err := r.redisUtil.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), string(sessionDataJSON), accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), string(sessionDataJSON), refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ToResponse(accessToken, refreshToken), http.StatusOK, nil
}

func (r *authService) Register(req requests.RegisterRequest) (resp responses.RegisterResponse, statusCode int, err error) {
	user := req.ReqToModel()

	hashedPassword, err := r.bcryptUtil.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	user.Password = string(hashedPassword)

	if err := r.transactionUtil.ExecuteInTransaction(func(tx *gorm.DB) error {
		if err = tx.Create(&user).Error; err != nil {
			if strings.Contains(err.Error(), "unique constraint") {
				return pkg.Error(errors.New("username already exists"))
			}

			return pkg.Error(err)
		}

		userHistory := models.UserHistory{
			Action:       constants.Action.Create,
			OldFirstName: user.FirstName,
			OldLastName:  user.LastName,
			OldBirthDate: user.BirthDate,
			OldPassword:  user.Password,
			OldEmail:     user.Email,
			OldVersion:   user.Version,
		}

		if err = tx.Create(&userHistory).Error; err != nil {
			return pkg.Error(err)
		}

		return nil
	}); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	fmt.Println(user.ID)

	return resp.ToResponse(user), http.StatusCreated, nil
}

func (r *authService) Logout(userContext contexts.UserContext) (statusCode int, err error) {
	if err := r.redisUtil.Del(context.Background(), fmt.Sprintf("access_token:%s", userContext.UserID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Del(context.Background(), fmt.Sprintf("refresh_token:%s", userContext.UserID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	return http.StatusOK, nil
}

func (r *authService) RefreshToken(userContext contexts.UserContext) (resp responses.LoginResponse, statusCode int, err error) {
	user := models.User{}
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

	sessionData := contexts.UserContext{
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

	if err := r.redisUtil.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), string(sessionDataJSON), accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), string(sessionDataJSON), refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ToResponse(accessToken, refreshToken), http.StatusOK, nil
}
