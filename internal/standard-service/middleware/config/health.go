package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func Healthz() fiber.Handler {
	return healthcheck.New(
		healthcheck.Config{
			LivenessEndpoint: "/healthz",
			LivenessProbe: func(c *fiber.Ctx) bool {
				return true
			},
			ReadinessEndpoint: "/readyz",
			ReadinessProbe: func(c *fiber.Ctx) bool {
				return true
			},
		},
	)
}
