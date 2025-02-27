package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/standard-golang/internal/configs"
	middlewareConfigs "github.com/wisaitas/standard-golang/internal/middlewares/configs"
	"github.com/wisaitas/standard-golang/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func init() {
	configs.LoadEnv()
}

type App struct {
	App     *fiber.App
	Configs *Configs
	routes  func()
}

func InitializeApp() *App {
	app := fiber.New()

	configs := InitializeConfigs()

	redisUtils := utils.NewRedisClient(configs.Redis)

	repositories := InitializeRepositories(configs.DB)
	services := InitializeServices(repositories, redisUtils)
	handlers := InitializeHandlers(services)
	validates := InitializeValidates()
	middlewares := InitializeMiddlewares(redisUtils)

	apiRoutes := app.Group("/api/v1")
	appRoutes := InitializeRoutes(apiRoutes, handlers, validates, middlewares)

	return &App{
		App:     app,
		Configs: configs,
		routes: func() {
			appRoutes.SetupRoutes()
		},
	}
}

func (r *App) SetupRoutes() {
	r.routes()
}

func (r *App) Run() {
	go func() {
		if err := r.App.Listen(fmt.Sprintf(":%s", configs.ENV.PORT)); err != nil {
			log.Fatalf("error starting server: %v\n", err)
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdown
	r.close()
}

func (r *App) close() {
	sqlDB, err := r.Configs.DB.DB()
	if err != nil {
		log.Fatalf("error getting database: %v\n", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", err)
	}

	if err := r.Configs.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", err)
	}

	log.Println("gracefully shutdown")
}

func (r *App) SetupMiddlewares() {
	r.App.Use(
		middlewareConfigs.Recovery(),
		middlewareConfigs.Limiter(),
		middlewareConfigs.CORS(),
		middlewareConfigs.Healthz(),
		middlewareConfigs.Logger(),
	)
}
