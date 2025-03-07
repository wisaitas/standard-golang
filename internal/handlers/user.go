package handlers

import (
	"errors"

	"github.com/wisaitas/standard-golang/internal/dtos/params"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/services/user"
	"github.com/wisaitas/standard-golang/internal/utils"

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
			Message: utils.Error(errors.New("failed to get queries")).Error(),
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
			Message: utils.Error(errors.New("failed to get request")).Error(),
		})
	}

	user, statusCode, err := r.userService.CreateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "user created successfully",
		Data:    user,
	})
}

func (r *UserHandler) UpdateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.UpdateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(errors.New("failed to get request")).Error(),
		})
	}

	param, ok := c.Locals("params").(params.UserParams)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(errors.New("failed to get params")).Error(),
		})
	}

	resp, statusCode, err := r.userService.UpdateUser(param, req)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "user updated successfully",
		Data:    resp,
	})
}
