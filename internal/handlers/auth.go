package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/models"
	authService "github.com/wisaitas/standard-golang/internal/services/auth"
)

type AuthHandler struct {
	authService authService.AuthService
}

func NewAuthHandler(authService authService.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (r *AuthHandler) Login(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.LoginRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: "failed to get request",
		})
	}

	resp, statusCode, err := r.authService.Login(req)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "login successfully",
		Data:    resp,
	})
}

func (r *AuthHandler) Register(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.RegisterRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: "failed to get request",
		})
	}

	resp, statusCode, err := r.authService.Register(req)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "user registered successfully",
		Data:    resp,
	})
}

func (r *AuthHandler) Logout(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: "user context not found",
		})
	}

	statusCode, err := r.authService.Logout(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "logout successfully",
	})
}

func (r *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: "user context not found",
		})
	}

	resp, statusCode, err := r.authService.RefreshToken(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "refresh token successfully",
		Data:    resp,
	})
}
