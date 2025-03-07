package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/models"
	authService "github.com/wisaitas/standard-golang/internal/services/auth"
	"github.com/wisaitas/standard-golang/internal/utils"
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
			Message: utils.Error(errors.New("failed to get request")).Error(),
		})
	}

	resp, statusCode, err := r.authService.Login(req)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
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
			Message: utils.Error(errors.New("failed to get request")).Error(),
		})
	}

	resp, statusCode, err := r.authService.Register(req)
	if err != nil {
		c.Locals("err", err)
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
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
			Message: utils.Error(errors.New("user context not found")).Error(),
		})
	}

	statusCode, err := r.authService.Logout(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
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
			Message: utils.Error(errors.New("user context not found")).Error(),
		})
	}

	resp, statusCode, err := r.authService.RefreshToken(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "refresh token successfully",
		Data:    resp,
	})
}
