package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

func (mw *MW) CheckAdminAccess(c *gin.Context) {
	var user database.User = c.MustGet("userData").(database.User)
	if !user.IsAdminUser {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Access Denied, Admin Access Required!"})
		c.Abort()
		return
	}
	c.Next()
}

func (mw *MW) CheckRootAccess(c *gin.Context) {
	var user database.User = c.MustGet("userData").(database.User)
	if user.Email != os.Getenv("ROOT_USER_EMAIL") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Access Denied, Root Access Required!"})
		c.Abort()
		return
	}
	c.Next()
}
