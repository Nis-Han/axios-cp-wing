package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/handlers"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/middleware"
)

func SetupRoutes(db *database.Queries) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/ping", handlers.Ping)

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/login", handlers.Login)
		userRoutes.POST("/signup", handlers.CreateUser)
	}

	authedRoutes := router.Group("/authed")
	authedRoutes.Use(middleware.AuthMiddleware)

	return router
}
