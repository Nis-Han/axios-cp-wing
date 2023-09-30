package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/handlers"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/login", handlers.Login)
		userRoutes.POST("/signup", handlers.CreateUser)
	}

	return router
}
