package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/request"
	"github.com/wisaitas/standard-golang/internal/dtos/response"
	"github.com/wisaitas/standard-golang/internal/services"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (r *AuthHandler) Login(c *fiber.Ctx) error {
	_, ok := c.Locals("req").(request.LoginRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get request",
		})
	}

	// user, statusCode, err := r.userService.Login(req)
	// if err != nil {
	// 	return c.Status(statusCode).JSON(response.ErrorResponse{
	// 		Message: err.Error(),
	// 	})
	// }

	return c.SendString("login")
}

func (r *AuthHandler) Register(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.RegisterRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "failed to get request",
		})
	}

	resp, statusCode, err := r.authService.Register(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(response.SuccessResponse{
		Message: "user registered successfully",
		Data:    resp,
	})
}
