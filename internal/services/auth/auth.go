package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wisaitas/standard-golang/internal/constants"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
	"github.com/wisaitas/standard-golang/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(req requests.LoginRequest) (resp responses.LoginResponse, statusCode int, err error)
	Register(req requests.RegisterRequest) (resp responses.RegisterResponse, statusCode int, err error)
	Logout(userContext models.UserContext) (statusCode int, err error)
	RefreshToken(userContext models.UserContext) (resp responses.LoginResponse, statusCode int, err error)
}

type authService struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	redis                 pkg.RedisClient
}

func NewAuthService(
	userRepository repositories.UserRepository,
	userHistoryRepository repositories.UserHistoryRepository,
	redis pkg.RedisClient,
) AuthService {
	return &authService{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		redis:                 redis,
	}
}

func (r *authService) Login(req requests.LoginRequest) (resp responses.LoginResponse, statusCode int, err error) {
	user := models.User{}
	if err := r.userRepository.GetBy(map[string]interface{}{"username": req.Username}, &user); err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, http.StatusNotFound, pkg.Error(err)
		}

		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return resp, http.StatusUnauthorized, pkg.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(time.Hour * 1)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	accessToken, err := utils.GenerateToken(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}, accessTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	refreshToken, err := utils.GenerateToken(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}, refreshTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), accessToken, accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), refreshToken, refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ToResponse(accessToken, refreshToken), statusCode, pkg.Error(err)
}

func (r *authService) Register(req requests.RegisterRequest) (resp responses.RegisterResponse, statusCode int, err error) {
	user := req.ReqToModel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	user.Password = string(hashedPassword)

	tx := r.userRepository.BeginTx()

	if err = tx.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			tx.Rollback()
			return resp, http.StatusBadRequest, pkg.Error(errors.New("username already exists"))
		}

		tx.Rollback()
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := tx.Create(
		&models.UserHistory{
			Action:      constants.ACTION.CREATE,
			UserID:      user.ID,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			BirthDate:   user.BirthDate,
			UserVersion: user.Version,
		},
	).Error; err != nil {
		tx.Rollback()
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ToResponse(user), http.StatusCreated, pkg.Error(err)
}

func (r *authService) Logout(userContext models.UserContext) (statusCode int, err error) {
	if err := r.redis.Del(context.Background(), fmt.Sprintf("access_token:%s", userContext.ID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Del(context.Background(), fmt.Sprintf("refresh_token:%s", userContext.ID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	return http.StatusOK, nil
}

func (r *authService) RefreshToken(userContext models.UserContext) (resp responses.LoginResponse, statusCode int, err error) {
	user := models.User{}
	if err := r.userRepository.GetBy(map[string]interface{}{"username": userContext.Username}, &user); err != nil {
		return resp, http.StatusNotFound, pkg.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(time.Hour * 1)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	accessToken, err := utils.GenerateToken(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}, accessTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	refreshToken, err := utils.GenerateToken(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	}, refreshTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), accessToken, accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), refreshToken, refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ToResponse(accessToken, refreshToken), http.StatusOK, nil
}
