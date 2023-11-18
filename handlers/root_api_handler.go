package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nerd500/axios-cp-wing/models"
)

func (api *Api) listAdmin(c *gin.Context) {
	adminList, err := api.DB.GetAllAdminUsers(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching admin users from DB"})
		return
	}

	c.JSON(http.StatusOK, models.DBAdminListtoAdminList(adminList))
}

func (api *Api) editAdminPermissions(c *gin.Context) {
	params := models.EditAdminAccessParams{}

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Malformed request Body, see readme.md for API Docs." + err.Error())})
		c.Abort()
		return
	}

	_, err := api.DB.GetUserFromEmail(c.Request.Context(), params.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("User with Email" + params.Email + "not found!" + err.Error())})
		c.Abort()
		return
	}

	user, err := api.DB.EditAdminAccess(c.Request.Context(), models.EditAdminAccessParamsToDBModel(params))

	if user.Email != params.Email || user.IsAdminUser != params.IsAdminUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Something Unexpected happened.\n DB Error:" + err.Error())})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("{is_admin_user:"+strconv.FormatBool(user.IsAdminUser)+"}"))
}
