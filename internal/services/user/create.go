package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type Create interface {
	CreateUser(req requests.CreateUserRequest) (resp responses.CreateUserResponse, statusCode int, err error)
}

type create struct {
	userRepository repositories.UserRepository
	redisUtil      utils.RedisClient
}

func NewCreate(
	userRepository repositories.UserRepository,
	redisUtil utils.RedisClient,
) Create {
	return &create{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *create) CreateUser(req requests.CreateUserRequest) (resp responses.CreateUserResponse, statusCode int, err error) {
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

	return resp.ToResponse(user), http.StatusCreated, nil
}
