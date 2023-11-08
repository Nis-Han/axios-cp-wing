package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/nerd500/axios-cp-wing/models"
	"github.com/nerd500/axios-cp-wing/utils"
	"github.com/stretchr/testify/assert"
)

var authDataList []models.AuthData
var loginCredentialsList []models.LoginData

func TestSignUp(t *testing.T) {
	newUser := models.User{
		FirstName: utils.GenerateRandomName(),
		LastName:  utils.GenerateRandomName(),
		Email:     utils.GenerateRandomEmail(),
		Password:  utils.GenerateRandomPassword(),
	}
	writer := makeRequest("POST", "/user/signup", newUser)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusCreated, writer.Code)

	authDataList = append(authDataList, models.AuthData{
		Email:     response["email"],
		AuthToken: response["auth"],
	})

	loginCredentialsList = append(loginCredentialsList, models.LoginData{
		Email:    newUser.Email,
		Password: newUser.Password,
	})
}

func TestLogin(t *testing.T) {
	loginCredentials := GetSampleLoginCredentials()
	writer := makeRequest("POST", "/user/login", loginCredentials)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]
	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}
