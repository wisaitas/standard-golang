package user

import (
	"errors"
	"net/http"
	"strings"

	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type Post interface {
	CreateUser(req request.CreateUserRequest) (resp response.CreateUserResponse, statusCode int, err error)
}

type post struct {
	userRepository repository.UserRepository
	redisUtil      redisPkg.Redis
}

func NewPost(
	userRepository repository.UserRepository,
	redisUtil redisPkg.Redis,
) Post {
	return &post{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *post) CreateUser(req request.CreateUserRequest) (resp response.CreateUserResponse, statusCode int, err error) {
	user := req.RequestToEntity()

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

	return resp.EntityToResponse(user), http.StatusCreated, nil
}
