package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/models"
	"github.com/nerd500/axios-cp-wing/utils"
)

func Login(c *gin.Context) {
	var loginData models.LoginData
	err := c.ShouldBindJSON(&loginData)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Ill-formatted request body")
		return
	}

	db, dbExists := c.Get("db")
	if !dbExists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Problem connecting with db"})
		return
	}

	databaseInstance, ok := db.(*database.Queries)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Problem connecting with db"})
		return
	}
	dbUser, err := databaseInstance.GetUser(c.Request.Context(), loginData.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	if !utils.CheckPassword(dbUser, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed: Wrong Password"})
		return
	}

	c.JSON(http.StatusOK, models.DbUserToUser(dbUser))

}

func CreateUser(c *gin.Context) {
	var newUser models.User
	var createUserParams database.CreateUserParams

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
	createUserParams.ID = uuid.New()
	createUserParams.Email = newUser.Email
	createUserParams.FirstName = newUser.FirstName
	createUserParams.LastName = newUser.LastName
	createUserParams.Salt = utils.GenerateSalt()
	createUserParams.HashedPassword = utils.HashPassword(newUser.Password, createUserParams.Salt)
	createUserParams.AuthToken = utils.GenerateAuthToken(100)
	createUserParams.IsAdminUser = false

	db, dbExists := c.Get("db")
	if !dbExists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Problem connecting with db"})
		return
	}

	databaseInstance, ok := db.(*database.Queries)

	_, err := databaseInstance.GetUser(c.Request.Context(), createUserParams.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	}

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Problem connecting with db"})
		return
	}

	var dbUser database.User
	dbUser, err = databaseInstance.CreateUser(c.Request.Context(), createUserParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt create user"})
		return
	}

	c.JSON(http.StatusCreated, models.DbUserToUser(dbUser))
}
