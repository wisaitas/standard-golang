package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/contexts"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/requests"
	authService "github.com/wisaitas/standard-golang/internal/standard-service/services/auth"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type authHandler struct {
	authService authService.AuthService
}

func NewAuthHandler(authService authService.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body requests.LoginRequest true "Login Request"
// @Success 200 {object} pkg.SuccessResponse
// @Router /auth/login [post]
func (r *authHandler) Login(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.LoginRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	resp, statusCode, err := r.authService.Login(req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "login successfully",
		Data:    resp,
	})
}

// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param registerRequest body requests.RegisterRequest true "Register Request"
// @Success 200 {object} pkg.SuccessResponse
// @Router /auth/register [post]
func (r *authHandler) Register(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.RegisterRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: errors.New("failed to get request").Error(),
		})
	}

	resp, statusCode, err := r.authService.Register(req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "user registered successfully",
		Data:    resp,
	})
}

// @Summary Logout
// @Description Logout
// @Tags Auth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} pkg.SuccessResponse
// @Router /auth/logout [post]
func (r *authHandler) Logout(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(contexts.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("user context not found")).Error(),
		})
	}

	statusCode, err := r.authService.Logout(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "logout successfully",
	})
}

// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} pkg.SuccessResponse
// @Router /auth/refresh-token [post]
func (r *authHandler) RefreshToken(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(contexts.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("user context not found")).Error(),
		})
	}

	resp, statusCode, err := r.authService.RefreshToken(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "refresh token successfully",
		Data:    resp,
	})
}
