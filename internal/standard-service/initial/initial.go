package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
)

func init() {
	if err := utils.ReadConfig(&env.Environment); err != nil {
		log.Fatalf("error reading config: %v\n", utils.Error(err))
	}
}

func InitializeApp() {
	clientConfig := newClientConfig()

	app := fiber.New()

	setupMiddleware(app)

	sharePkg := newSharePkg(clientConfig)

	repository := newRepository(clientConfig)
	service := newService(repository, sharePkg)
	handler := newHandler(service)
	validate := newValidate(sharePkg)
	middleware := newMiddleware(sharePkg)

	newRoute(app, handler, validate, middleware)

	run(app, clientConfig)
}

func run(app *fiber.App, clientConfig *clientConfig) {
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", env.Environment.Server.Port)); err != nil {
			log.Fatalf("error starting server: %v\n", utils.Error(err))
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdown

	close(app, clientConfig)
}

func close(app *fiber.App, clientConfig *clientConfig) {
	sqlDB, err := clientConfig.DB.DB()
	if err != nil {
		log.Fatalf("error getting database: %v\n", utils.Error(err))
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", utils.Error(err))
	}

	if err := clientConfig.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", utils.Error(err))
	}

	if err := app.Shutdown(); err != nil {
		log.Fatalf("error shutting down app: %v\n", utils.Error(err))
	}

	log.Println("gracefully shutdown")
}
