package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/models"
)

func AuthMiddleware(c *gin.Context) {
	var authData models.AuthData
	err := c.ShouldBindJSON(&authData)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Ill-formatted request body")
		return
	}

	var userAuthTokenwithEmailParams = database.GetUserAuthTokenwithEmailParams{AuthToken: authData.AuthToken, Email: authData.Email}
	userData, err := database.DBInstance.GetUserAuthTokenwithEmail(c.Request.Context(), userAuthTokenwithEmailParams)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Invalid AuthToken for user Email: %v", authData.Email)})
		return
	}

	c.Set("userData", userData)

	c.Next()
}
