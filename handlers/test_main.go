package handlers

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/middleware"
)

func router() *gin.Engine {
	router := gin.Default()

	// HealthCheck Route
	router.GET("/ping", Ping)

	// User Auth Routes
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/login", Login)
		userRoutes.POST("/signup", CreateUser)
	}

	// Authed User Routes
	authedRoutes := router.Group("/authed")
	authedRoutes.Use(middleware.AuthMiddleware)

	// Admin Routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware)
	adminRoutes.Use(func(c *gin.Context) {
		var user database.User = c.MustGet("userData").(database.User)
		if !user.IsAdminUser {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Access Denied"})
		}
		c.Next()
	})
	{
		adminRoutes.POST("/createTask", CreateTask)
	}

	return router
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	setupServer()
	defer database.CloseDataBase()

	exitCode := m.Run()

	os.Exit(exitCode)
}

func setTestEnv() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initialiseTestDB() {
	if err := database.InitialiseDatabase(); err != nil {
		log.Fatal(err)
	}
}

func setupServer() {

	setTestEnv()

	initialiseTestDB()

}
