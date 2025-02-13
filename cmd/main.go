package main

import (
	"flag"
	"fmt"
	"go-echo-experiment/internal/controller"
	"go-echo-experiment/internal/repository"
	"go-echo-experiment/internal/routes"
	"go-echo-experiment/internal/service"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go-echo-experiment/config"
	"go-echo-experiment/pkg/middleware"
)

func main() {
	// flag --migrate
	migrate := flag.Bool("migrate", false, "Migrate database")

	flag.Parse()

	if *migrate {
		fmt.Println("Running database migration...")
		config.RunMigration()
		return
	}

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectGORM()

	// Initialize Echo
	e := echo.New()

	e.Use(middleware.Logger)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := controller.NewUserHandler(userService)

	// Connect to database
	//db := config.ConnectDB()

	// Register routes
	routesGroup := e.Group("/v1")
	routes.UserRoutes(routesGroup, userHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo API is running")
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
