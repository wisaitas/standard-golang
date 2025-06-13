package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"github.com/wisaitas/standard-golang/pkg"
)

func init() {
	if err := pkg.ReadConfig(&env.Environment); err != nil {
		log.Fatalf("error reading config: %v\n", pkg.Error(err))
	}
}

func InitializeApp() {
	clientConfig := newClientConfig()

	app := fiber.New()

	setupMiddleware(app)

	lib := newLib(clientConfig)

	repository := newRepository(clientConfig)
	service := newService(repository, lib)
	handler := newHandler(service)
	validate := newValidate(lib)
	middleware := newMiddleware(lib)

	newRoute(app, handler, validate, middleware)

	run(app, clientConfig)
}

func run(app *fiber.App, clientConfig *clientConfig) {
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", env.Environment.Server.Port)); err != nil {
			log.Fatalf("error starting server: %v\n", pkg.Error(err))
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
		log.Fatalf("error getting database: %v\n", pkg.Error(err))
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", pkg.Error(err))
	}

	if err := clientConfig.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", pkg.Error(err))
	}

	if err := app.Shutdown(); err != nil {
		log.Fatalf("error shutting down app: %v\n", pkg.Error(err))
	}

	log.Println("gracefully shutdown")
}
