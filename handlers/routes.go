package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/middleware"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// HealthCheck API
	router.GET("/ping", Ping)

	// Root User API
	rootUserRoutes := router.Group("/root")
	rootUserRoutes.Use(middleware.AuthMiddleware)
	rootUserRoutes.Use(middleware.CheckRootAccess)
	{
		rootUserRoutes.GET("/listAdmin", listAdmin)
		rootUserRoutes.GET("/listUser", listUser)
		rootUserRoutes.PATCH("/updateAdminPermission", editAdminPermissions)
	}

	// User Login/SIgnup APIs
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/login", Login)
		userRoutes.POST("/signup", CreateUser)
	}

	// User level Authed APIs
	authedRoutes := router.Group("/api")
	authedRoutes.Use(middleware.AuthMiddleware)

	// Admin Level Authed APIs
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware)
	adminRoutes.Use(middleware.CheckAdminAccess)
	{
		adminRoutes.POST("/createTask", CreateTask)
		adminRoutes.GET("/tasks", GetAllTasks)
	}
	return router
}
