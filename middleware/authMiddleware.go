package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

type MW struct {
	DB database.Querier
}

func (mw *MW) AuthMiddlewareForVerifiedEmail(c *gin.Context) {
	authData := c.Request.Header.Get("AuthToken")

	userData, err := mw.DB.GetUserFromAuthToken(c.Request.Context(), authData)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid AuthToken"})
		c.Abort()
		return
	}

	if !userData.VerifiedUser {
		c.JSON(http.StatusLocked, gin.H{"message": "User Need to verify their Emaail before accessing this endpoint"})
		c.Abort()
		return
	}

	c.Set("userData", userData)

	c.Next()
}

func (mw *MW) AuthMiddlewareForUnverifiedEmail(c *gin.Context) {
	authData := c.Request.Header.Get("AuthToken")

	userData, err := mw.DB.GetUserFromAuthToken(c.Request.Context(), authData)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid AuthToken"})
		c.Abort()
		return
	}

	c.Set("userData", userData)

	c.Next()
}
