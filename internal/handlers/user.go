package handlers

import (
	"github.com/wisaitas/standard-golang/internal/dtos/request"
	"github.com/wisaitas/standard-golang/internal/dtos/response"
	"github.com/wisaitas/standard-golang/internal/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(
	userService services.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (r *UserHandler) GetUsers(c *fiber.Ctx) error {
	querys, ok := c.Locals("querys").(request.PaginationParam)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get querys",
		})
	}

	users, statusCode, err := r.userService.GetUsers(querys)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "users fetched successfully",
		Data:    users,
	})
}

func (r *UserHandler) CreateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.CreateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get request",
		})
	}

	user, statusCode, err := r.userService.CreateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "user created successfully",
		Data:    user,
	})
}
