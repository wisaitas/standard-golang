package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	contextPkg "github.com/wisaitas/share-pkg/auth/context"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	authService "github.com/wisaitas/standard-golang/internal/standard-service/service/auth"
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

func (r *authHandler) Login(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.LoginRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get request")),
		})
	}

	resp, statusCode, err := r.authService.Login(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(response.ApiResponse[any]{
		Message: "login successfully",
		Data:    resp,
	})
}

func (r *authHandler) Register(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(request.RegisterRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get request")),
		})
	}

	resp, statusCode, err := r.authService.Register(req)
	if err != nil {
		return c.Status(statusCode).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(response.ApiResponse[any]{
		Message: "user registered successfully",
		Data:    resp,
	})
}

func (r *authHandler) Logout(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(contextPkg.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse[any]{
			Error: utils.Error(errors.New("user context not found")),
		})
	}

	statusCode, err := r.authService.Logout(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(response.ApiResponse[any]{
		Message: "logout successfully",
	})
}

func (r *authHandler) RefreshToken(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(contextPkg.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse[any]{
			Error: utils.Error(errors.New("user context not found")),
		})
	}

	resp, statusCode, err := r.authService.RefreshToken(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(response.ApiResponse[any]{
		Message: "refresh token successfully",
		Data:    resp,
	})
}
