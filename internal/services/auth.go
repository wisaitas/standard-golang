package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/wisaitas/standard-golang/internal/dtos/request"
	"github.com/wisaitas/standard-golang/internal/dtos/response"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(req request.LoginRequest) (resp response.LoginResponse, statusCode int, err error)
	Register(req request.RegisterRequest) (resp response.RegisterResponse, statusCode int, err error)
	Logout(userContext models.UserContext) (statusCode int, err error)
	RefreshToken(userContext models.UserContext) (resp response.LoginResponse, statusCode int, err error)
}

type authService struct {
	userRepository repositories.UserRepository
	redis          utils.RedisClient
}

func NewAuthService(
	userRepository repositories.UserRepository,
	redis utils.RedisClient,
) AuthService {
	return &authService{
		userRepository: userRepository,
		redis:          redis,
	}
}

func (r *authService) Login(req request.LoginRequest) (resp response.LoginResponse, statusCode int, err error) {
	user := models.User{}
	if err := r.userRepository.GetBy("username", req.Username, &user); err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, http.StatusNotFound, err
		}

		return resp, http.StatusInternalServerError, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return resp, http.StatusUnauthorized, err
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
		return resp, http.StatusInternalServerError, err
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
		return resp, http.StatusInternalServerError, err
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), refreshToken, refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return resp.ToResponse(accessToken, refreshToken), statusCode, err
}

func (r *authService) Register(req request.RegisterRequest) (resp response.RegisterResponse, statusCode int, err error) {
	user := req.ToModel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	user.Password = string(hashedPassword)

	if err = r.userRepository.Create(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return resp, http.StatusBadRequest, errors.New("username already exists")
		}

		return resp, http.StatusInternalServerError, err
	}

	return resp.ToResponse(user), http.StatusCreated, err
}

func (r *authService) Logout(userContext models.UserContext) (statusCode int, err error) {
	if err := r.redis.Del(context.Background(), fmt.Sprintf("access_token:%s", userContext.ID)); err != nil {
		return http.StatusInternalServerError, err
	}

	if err := r.redis.Del(context.Background(), fmt.Sprintf("refresh_token:%s", userContext.ID)); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (r *authService) RefreshToken(userContext models.UserContext) (resp response.LoginResponse, statusCode int, err error) {
	user := models.User{}
	if err := r.userRepository.GetBy("username", userContext.Username, &user); err != nil {
		return resp, http.StatusNotFound, err
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
		return resp, http.StatusInternalServerError, err
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
		return resp, http.StatusInternalServerError, err
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), refreshToken, refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, err
	}

	return resp.ToResponse(accessToken, refreshToken), http.StatusOK, nil
}
