package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/standard-golang/internal/configs"
	"github.com/wisaitas/standard-golang/internal/handlers"
	middlewareConfigs "github.com/wisaitas/standard-golang/internal/middlewares/configs"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/routes"
	"github.com/wisaitas/standard-golang/internal/services"
	"github.com/wisaitas/standard-golang/internal/validates"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func init() {
	configs.LoadEnv()
}

type App struct {
	App    *fiber.App
	DB     *gorm.DB
	routes func()
}

func InitializeApp() *App {
	app := fiber.New()
	db := configs.ConnectDB()

	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userRepository)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize validates
	userValidate := validates.NewUserValidate()
	authValidate := validates.NewAuthValidate()

	// Initialize routes
	apiRoutes := app.Group("/api/v1")
	userRoutes := routes.NewUserRoutes(apiRoutes, userHandler, userValidate)
	authRoutes := routes.NewAuthRoutes(apiRoutes, authHandler, authValidate)

	return &App{
		App: app,
		DB:  db,
		routes: func() {
			userRoutes.UserRoutes()
			authRoutes.AuthRoutes()
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
	sqlDB, err := r.DB.DB()
	if err != nil {
		log.Fatalf("error getting database: %v\n", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", err)
	}
}

func (r *App) SetupMiddlewares() {
	r.App.Use(
		middlewareConfigs.Limiter(),
		middlewareConfigs.CORS(),
		middlewareConfigs.Healthz(),
		middlewareConfigs.Logger(),
	)
}
