package config

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/utils"
)

func Recovery() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			c.Status(fiber.StatusInternalServerError).JSON(response.ApiResponse[any]{
				Error: utils.Error(errors.New("internal server error")),
			})
		},
	})
}
