package configs

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/utils"
)

func Recovery() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			c.Status(fiber.StatusInternalServerError).JSON(responses.ErrorResponse{
				Message: utils.Error(errors.New("internal server error")).Error(),
			})
		},
	})
}
