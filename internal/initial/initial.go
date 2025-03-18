package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/standard-golang/internal/configs"
	middlewareConfigs "github.com/wisaitas/standard-golang/internal/middlewares/configs"
	"github.com/wisaitas/standard-golang/pkg"

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

	configs := initializeConfigs()

	redisUtils := pkg.NewRedisClient(configs.Redis)

	repositories := initializeRepositories(configs.DB)
	services := initializeServices(repositories, redisUtils)
	handlers := initializeHandlers(services)
	validates := initializeValidates()
	middlewares := initializeMiddlewares(redisUtils)

	apiRoutes := app.Group("/api/v1")
	appRoutes := initializeRoutes(apiRoutes, handlers, validates, middlewares)

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
			log.Fatalf("error starting server: %v\n", pkg.Error(err))
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
		log.Fatalf("error getting database: %v\n", pkg.Error(err))
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", pkg.Error(err))
	}

	if err := r.Configs.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", pkg.Error(err))
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
