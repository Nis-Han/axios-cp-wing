package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nerd500/axios-cp-wing/models"
	"github.com/nerd500/axios-cp-wing/utils"
)

var users []models.User

func Login(c *gin.Context) {
	var loginData models.LoginData
	err := c.ShouldBindJSON(&loginData)

	fmt.Println(loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Ill-formatted request body")
		return
	}

	// TODO: Implement password verification logic once DB is connected
	var DBUser models.User
	for _, user := range users {
		if user.Email == loginData.Email {
			DBUser = user
		}
	}

	if !utils.CheckPassword(DBUser, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, DBUser)

}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.IsValidEmail(newUser.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.IsValidPassword(newUser.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.ID = uuid.New().String()
	newUser.Salt = utils.GenerateSalt()
	newUser.Password = utils.HashPassword(newUser.Password, newUser.Salt)
	newUser.Auth = utils.GenerateAuthToken(100)

	// TODO: connect with DB and add logic
	users = append(users, newUser)

	fmt.Println(newUser)
	c.JSON(http.StatusCreated, newUser)
}
