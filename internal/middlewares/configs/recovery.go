package middleware_configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
)

func Recovery() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: "Internal Server Error",
			})
		},
	})
}
