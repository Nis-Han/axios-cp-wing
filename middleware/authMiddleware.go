package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

type MW struct {
	DB database.Querier
}

func (mw *MW) AuthMiddleware(c *gin.Context) {
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
