package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
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

	exitCode := m.Run()

	os.Exit(exitCode)
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	log.Print(bytes.NewBuffer(requestBody))
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	writer := httptest.NewRecorder()
	router().ServeHTTP(writer, request)
	return writer
}
