package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func Pprof() fiber.Handler {
	return pprof.New()
}
