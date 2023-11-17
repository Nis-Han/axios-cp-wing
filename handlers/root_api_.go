package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/models"
)

func listAdmin(c *gin.Context) {
	adminList, err := database.DBInstance.GetAllAdminUsers(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching admin users from DB"})
		return
	}

	c.JSON(http.StatusOK, models.DBAdminListtoAdminList(adminList))
}

func listUser(c *gin.Context) {
	userList, err := database.DBInstance.GetAllUsers(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching users from DB"})
		return
	}

	c.JSON(http.StatusOK, models.DBUserListtoUserList(userList))
}
