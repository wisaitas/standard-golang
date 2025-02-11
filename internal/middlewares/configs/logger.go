package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() fiber.Handler {
	return logger.New(
		logger.Config{
			Format: "${ua} - [${ip}:${port}] - ${status} - ${method} - ${path}\n",
			Done: func(c *fiber.Ctx, logString []byte) {
				// if string(c.Request().Header.ContentType()) == "application/json" {
				// 	fmt.Printf("request : %s\n", string(c.Request().Body()))
				// }

				// if string(c.Response().Header.ContentType()) == "application/json" {
				// 	var prettyJSON bytes.Buffer
				// 	if err := json.Indent(&prettyJSON, c.Response().Body(), "", "    "); err == nil {
				// 		fmt.Printf("response\n %s\n", prettyJSON.String())
				// 	} else {
				// 		fmt.Printf("response\n %s\n", string(c.Response().Body()))
				// 	}
				// }
			},
		},
	)
}
