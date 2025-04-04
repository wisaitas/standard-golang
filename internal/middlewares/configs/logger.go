package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() fiber.Handler {
	return logger.New(
		logger.Config{
			Format: fmt.Sprintf("[%s ${time}]--[${ua}]--[${ip}:${port}]--[${status}]--[${method}]--[${path}]\n", time.Now().Format("2006-01-02")),
			Done: func(c *fiber.Ctx, logString []byte) {
				if c.Response().StatusCode() != 200 && c.Response().StatusCode() != 201 && c.Response().StatusCode() != 204 {
					if string(c.Request().Header.ContentType()) == "application/json" {
						requestBody := string(c.Request().Body())
						compactRequest := new(bytes.Buffer)
						if err := json.Compact(compactRequest, []byte(requestBody)); err == nil {
							log.Printf("request: %s\n", compactRequest.String())
						} else {
							log.Printf("request: %s\n", requestBody)
						}
					}
				}
			},
		},
	)
}
