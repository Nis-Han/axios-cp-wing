package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/middleware"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Ping)
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/login", Login)
		userRoutes.POST("/signup", CreateUser)
	}
	authedRoutes := router.Group("/authed")
	authedRoutes.Use(middleware.AuthMiddleware)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware)
	adminRoutes.Use(func(c *gin.Context) {
		var user database.User = c.MustGet("userData").(database.User)
		if !user.IsAdminUser {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Access Denied"})
			c.Abort()
			return
		}
		c.Next()
	})
	{
		adminRoutes.POST("/createTask", CreateTask)
		adminRoutes.GET("/tasks", GetAllTasks)
	}
	return router
}
