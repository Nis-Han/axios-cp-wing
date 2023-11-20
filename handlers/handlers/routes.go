package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/client/email_client"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/middleware"
)

type Api struct {
	DB          database.Querier
	EmailClient email_client.EmailClientInterface
}

func SetupRoutes(api *Api) *gin.Engine {
	router := gin.Default()
	mw := middleware.MW{DB: api.DB}

	// HealthCheck API
	router.GET("/ping", api.Ping)

	// Root User API
	rootUserRoutes := router.Group("/root")
	rootUserRoutes.Use(mw.AuthMiddlewareForVerifiedEmail)
	rootUserRoutes.Use(mw.CheckRootAccess)
	{
		rootUserRoutes.GET("/listAdmin", api.listAdmin)
		rootUserRoutes.PATCH("/updateAdminPermission", api.editAdminPermissions)
	}

	// User Login/SIgnup APIs
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/login", api.Login)
		userRoutes.POST("/signup", api.CreateUser)
	}

	// OTP Validation Routes
	otpRoutes := router.Group("/verify")
	otpRoutes.Use(mw.AuthMiddlewareForUnverifiedEmail)
	{
		otpRoutes.POST("/createOTP", api.generateAndSendOTPViaEmail)
		otpRoutes.POST("/verifyOTP", api.verifyUserFromOTP)
	}

	// User level Authed APIs
	authedRoutes := router.Group("/api")
	authedRoutes.Use(mw.AuthMiddlewareForVerifiedEmail)

	// Admin Level Authed APIs
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(mw.AuthMiddlewareForVerifiedEmail)
	adminRoutes.Use(mw.CheckAdminAccess)
	{
		adminRoutes.GET("/listUser", api.listUser)
		adminRoutes.POST("/createTask", api.CreateTask)
		adminRoutes.GET("/tasks", api.GetAllTasks)
	}
	return router
}
