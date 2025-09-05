package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/utils"
	standardservice "github.com/wisaitas/standard-golang/internal/standard-service"
)

func init() {
	if err := utils.ReadConfig(&standardservice.ENV); err != nil {
		log.Fatalf("error reading config: %v\n", utils.Error(err))
	}
}

type App struct {
	App          *fiber.App
	ClientConfig *clientConfig
}

func InitializeApp() *App {
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

	return &App{
		App:          app,
		ClientConfig: clientConfig,
	}
}

func (a *App) Run() chan os.Signal {
	go func() {
		if err := a.App.Listen(fmt.Sprintf(":%d", standardservice.ENV.Server.Port)); err != nil {
			log.Fatalf("error starting server: %v\n", utils.Error(err))
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdown

	return gracefulShutdown
}

func (a *App) Cleanup() {
	sqlDB, err := a.ClientConfig.DB.DB()
	if err != nil {
		log.Fatalf("error getting database: %v\n", utils.Error(err))
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", utils.Error(err))
	}

	if err := a.ClientConfig.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", utils.Error(err))
	}

	if err := a.App.Shutdown(); err != nil {
		log.Fatalf("error shutting down app: %v\n", utils.Error(err))
	}

	log.Println("gracefully shutdown")
}
