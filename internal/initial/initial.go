package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/standard-golang/internal/env"
	"github.com/wisaitas/standard-golang/pkg"

	"github.com/gofiber/fiber/v2"
)

func init() {
	env.LoadEnv()
}

func InitializeApp() {
	config := newConfig()

	app := fiber.New()

	util := newUtil(config)

	repository := newRepository(config)
	service := newService(repository, util)
	handler := newHandler(service)
	validate := newValidate(util)
	middleware := newMiddleware(util)

	newRoute(app, handler, validate, middleware)

	run(app, config)
}

func run(app *fiber.App, configs *config) {
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", env.PORT)); err != nil {
			log.Fatalf("error starting server: %v\n", pkg.Error(err))
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdown

	close(app, configs)
}

func close(app *fiber.App, config *config) {
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatalf("error getting database: %v\n", pkg.Error(err))
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", pkg.Error(err))
	}

	if err := config.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", pkg.Error(err))
	}

	if err := app.Shutdown(); err != nil {
		log.Fatalf("error shutting down app: %v\n", pkg.Error(err))
	}

	log.Println("gracefully shutdown")
}
