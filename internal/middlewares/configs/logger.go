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

						var requestMap map[string]interface{}
						if err := json.Unmarshal([]byte(requestBody), &requestMap); err == nil {
							if _, exists := requestMap["password"]; exists {
								requestMap["password"] = "*********"
							}

							if _, exists := requestMap["confirm_password"]; exists {
								requestMap["confirm_password"] = "*********"
							}

							if maskedJSON, err := json.Marshal(requestMap); err == nil {
								compactRequest := new(bytes.Buffer)
								if err := json.Compact(compactRequest, maskedJSON); err == nil {
									log.Printf("request: %s\n", compactRequest.String())
								}
							}
						} else {
							log.Printf("request: %s\n", requestBody)
						}
					}
				}
			},
		},
	)
}
