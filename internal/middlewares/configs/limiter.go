package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Limiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,             // 100 request per minute same ip
		Expiration: 1 * time.Minute, // 1 minute reset counter
	})
}
