package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/standard-service/repositories"
	"github.com/wisaitas/standard-golang/pkg"
	"golang.org/x/crypto/bcrypt"
)

type Post interface {
	CreateUser(req requests.CreateUserRequest) (resp responses.CreateUserResponse, statusCode int, err error)
}

type post struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisUtil
}

func NewPost(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisUtil,
) Post {
	return &post{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *post) CreateUser(req requests.CreateUserRequest) (resp responses.CreateUserResponse, statusCode int, err error) {
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

	return resp.ModelToResponse(user), http.StatusCreated, nil
}
