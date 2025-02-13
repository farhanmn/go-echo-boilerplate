package routes

import (
	"github.com/labstack/echo/v4"
	"go-echo-experiment/internal/controller"
)

func UserRoutes(g *echo.Group, userHandler *controller.UserHandler) {
	r := g.Group("/auth")
	
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/:email", userHandler.GetUser)
}
