package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}
