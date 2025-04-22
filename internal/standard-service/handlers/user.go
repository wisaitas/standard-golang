package handlers

import (
	"errors"

	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/params"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/standard-service/services/user"
	"github.com/wisaitas/standard-golang/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(
	userService user.UserService,
) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// @Summary Get Users
// @Description Get Users
// @Tags User
// @Accept json
// @Produce json
// @Param query query pkg.PaginationQuery true "Pagination Query"
// @Success 200 {object} pkg.SuccessResponse
// @Router /users [get]
func (r *userHandler) GetUsers(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(pkg.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	users, statusCode, err := r.userService.GetUsers(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	if len(users) == 0 {
		users = []responses.GetUsersResponse{}
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "users fetched successfully",
		Data:    users,
	})
}

func (r *userHandler) CreateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.CreateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	user, statusCode, err := r.userService.CreateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "user created successfully",
		Data:    user,
	})
}

// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param params path params.UserParams true "User Params"
// @Success 200 {object} pkg.SuccessResponse
// @Router /users/{id} [put]
func (r *userHandler) UpdateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.UpdateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	param, ok := c.Locals("params").(params.UserParams)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get params")).Error(),
		})
	}

	resp, statusCode, err := r.userService.UpdateUser(param, req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "user updated successfully",
		Data:    resp,
	})
}
