package handlers

import (
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/services/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(
	userService user.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (r *UserHandler) GetUsers(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: "failed to get queries",
		})
	}

	users, statusCode, err := r.userService.GetUsers(query)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "users fetched successfully",
		Data:    users,
	})
}

func (r *UserHandler) CreateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.CreateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: "failed to get request",
		})
	}

	user, statusCode, err := r.userService.CreateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "user created successfully",
		Data:    user,
	})
}
