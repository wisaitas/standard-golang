package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	contextPkg "github.com/wisaitas/share-pkg/auth/context"
	"github.com/wisaitas/share-pkg/auth/jwt"
	"github.com/wisaitas/share-pkg/cache/redis"
	bcryptPkg "github.com/wisaitas/share-pkg/crypto/bcrypt"
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	transactionmanager "github.com/wisaitas/share-pkg/db/transaction-manager"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/constant"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(req request.LoginRequest) (resp response.LoginResponse, statusCode int, err error)
	Register(req request.RegisterRequest) (resp response.RegisterResponse, statusCode int, err error)
	Logout(userContext contextPkg.UserContext) (statusCode int, err error)
	RefreshToken(userContext contextPkg.UserContext) (resp response.LoginResponse, statusCode int, err error)
}

type authService struct {
	userRepository        repository.UserRepository
	userHistoryRepository repository.UserHistoryRepository
	redis                 redis.Redis
	bcrypt                bcryptPkg.Bcrypt
	jwt                   jwt.Jwt
	transactionManager    transactionmanager.TransactionManager
}

func NewAuthService(
	userRepository repository.UserRepository,
	userHistoryRepository repository.UserHistoryRepository,
	redis redis.Redis,
	bcrypt bcryptPkg.Bcrypt,
	jwt jwt.Jwt,
	transactionManager transactionmanager.TransactionManager,
) AuthService {
	return &authService{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		redis:                 redis,
		bcrypt:                bcrypt,
		jwt:                   jwt,
		transactionManager:    transactionManager,
	}
}

func (r *authService) Login(req request.LoginRequest) (resp response.LoginResponse, statusCode int, err error) {
	user := entity.User{}

	if err := r.userRepository.GetBy(&user, repositoryPkg.NewCondition("username = ?", req.Username), nil); err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, http.StatusNotFound, utils.Error(err)
		}

		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return resp, http.StatusUnauthorized, utils.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(time.Hour * 1)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	tokenData := map[string]interface{}{
		"user_id": user.Id,
	}

	accessToken, err := r.jwt.GenerateToken(tokenData, accessTokenExp.Unix(), env.Environment.Server.JwtSecret)
	if err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	refreshToken, err := r.jwt.GenerateToken(tokenData, refreshTokenExp.Unix(), env.Environment.Server.JwtSecret)
	if err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	sessionData := contextPkg.UserContext{
		UserID:       user.Id,
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
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.Id), string(sessionDataJSON), accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.Id), string(sessionDataJSON), refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	return resp.EntityToResponse(accessToken, refreshToken), http.StatusOK, nil
}

func (r *authService) Register(req request.RegisterRequest) (resp response.RegisterResponse, statusCode int, err error) {
	user := req.RequestToEntity()

	hashedPassword, err := r.bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	user.Password = string(hashedPassword)

	if err := r.transactionManager.ExecuteInTransaction(func(tx *gorm.DB) error {
		txUserRepository := r.userRepository.WithTx(tx)
		txUserHistoryRepository := r.userHistoryRepository.WithTx(tx)

		if err := txUserRepository.Create(&user); err != nil {
			return err
		}

		userHistory := entity.UserHistory{
			Action:       constant.Action.Create,
			OldFirstName: user.FirstName,
			OldLastName:  user.LastName,
			OldBirthDate: user.BirthDate,
			OldPassword:  user.Password,
			OldEmail:     user.Email,
			OldVersion:   user.Version,
		}

		if err := txUserHistoryRepository.Create(&userHistory); err != nil {
			return err
		}

		return nil
	}); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return resp, http.StatusBadRequest, utils.Error(errors.New("username already exists"))
		}

		return resp, http.StatusInternalServerError, utils.Error(err)
	}
	return resp.EntityToResponse(user), http.StatusCreated, nil
}

func (r *authService) Logout(userContext contextPkg.UserContext) (statusCode int, err error) {
	if err := r.redis.Del(context.Background(), fmt.Sprintf("access_token:%s", userContext.UserID)); err != nil {
		return http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redis.Del(context.Background(), fmt.Sprintf("refresh_token:%s", userContext.UserID)); err != nil {
		return http.StatusInternalServerError, utils.Error(err)
	}

	return http.StatusOK, nil
}

func (r *authService) RefreshToken(userContext contextPkg.UserContext) (resp response.LoginResponse, statusCode int, err error) {
	user := entity.User{}
	if err := r.userRepository.GetBy(&user, repositoryPkg.NewCondition("username = ?", userContext.Username), nil); err != nil {
		return resp, http.StatusNotFound, utils.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(time.Hour * 1)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	tokenData := map[string]interface{}{
		"user_id": user.Id,
	}

	accessToken, err := r.jwt.GenerateToken(tokenData, accessTokenExp.Unix(), env.Environment.Server.JwtSecret)
	if err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	refreshToken, err := r.jwt.GenerateToken(tokenData, refreshTokenExp.Unix(), env.Environment.Server.JwtSecret)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	sessionData := contextPkg.UserContext{
		UserID:       user.Id,
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
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.Id), string(sessionDataJSON), accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.Id), string(sessionDataJSON), refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	return resp.EntityToResponse(accessToken, refreshToken), http.StatusOK, nil
}
