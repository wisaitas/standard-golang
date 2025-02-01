package validates

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func validateCommonRequest[T any](c *fiber.Ctx, req *T) error {
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := validator.New().Struct(req); err != nil {
		return err
	}

	return nil
}
