package services

import (
	"errors"
	"net/http"
	"strings"

	"github.com/wisaitas/standard-golang/internal/dtos/request"
	"github.com/wisaitas/standard-golang/internal/dtos/response"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers(querys request.PaginationParam) (resp []response.GetUsersResponse, statusCode int, err error)
	CreateUser(req request.CreateUserRequest) (resp response.CreateUserResponse, statusCode int, err error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (r *userService) GetUsers(querys request.PaginationParam) (resp []response.GetUsersResponse, statusCode int, err error) {
	users := []models.User{}

	if err := r.userRepository.GetAll(&users, &querys); err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, err
	}

	for _, user := range users {
		respGetUser := response.GetUsersResponse{}
		resp = append(resp, respGetUser.ToResponse(user))
	}

	return resp, http.StatusOK, nil

}

func (r *userService) CreateUser(req request.CreateUserRequest) (resp response.CreateUserResponse, statusCode int, err error) {
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
