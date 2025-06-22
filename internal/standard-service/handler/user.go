package handler

import (
	"errors"

	responsePkg "github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/param"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	userService "github.com/wisaitas/standard-golang/internal/standard-service/service/user"

	"github.com/gofiber/fiber/v2"
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/share-pkg/utils"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type userHandler struct {
	userService userService.UserService
}

func NewUserHandler(
	userService userService.UserService,
) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (r *userHandler) GetUsers(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(repositoryPkg.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responsePkg.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get queries")),
		})
	}

	users, statusCode, err := r.userService.GetUsers(query)
	if err != nil {
		return c.Status(statusCode).JSON(responsePkg.ApiResponse[any]{
			Error: err,
		})
	}

	if len(users) == 0 {
		users = []response.GetUsersResponse{}
	}

	return c.Status(statusCode).JSON(responsePkg.ApiResponse[[]response.GetUsersResponse]{
		Message: "users fetched successfully",
		Data:    users,
	})
}

func (r *userHandler) CreateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.CreateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responsePkg.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get request")),
		})
	}

	user, statusCode, err := r.userService.CreateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(responsePkg.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(responsePkg.ApiResponse[response.CreateUserResponse]{
		Message: "user created successfully",
		Data:    user,
	})
}

func (r *userHandler) UpdateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.UpdateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responsePkg.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get request")),
		})
	}

	param, ok := c.Locals("params").(param.UserParam)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responsePkg.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get params")),
		})
	}

	resp, statusCode, err := r.userService.UpdateUser(param, req)
	if err != nil {
		return c.Status(statusCode).JSON(responsePkg.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(responsePkg.ApiResponse[response.UpdateUserResponse]{
		Message: "user updated successfully",
		Data:    resp,
	})
}
