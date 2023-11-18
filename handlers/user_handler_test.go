package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/nerd500/axios-cp-wing/internal/database"
	mockdb "github.com/nerd500/axios-cp-wing/internal/database/mock"
	"github.com/nerd500/axios-cp-wing/models"
	"github.com/nerd500/axios-cp-wing/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSignUp(t *testing.T) {
	newUser := models.User{
		FirstName: utils.GenerateRandomName(),
		LastName:  utils.GenerateRandomName(),
		Email:     utils.GenerateRandomEmail(),
		Password:  utils.GenerateRandomPassword(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromEmail(gomock.Any(), gomock.Eq(newUser.Email)).
		Times(1).
		Return(database.User{}, errors.New("User Not Found Error"))

	MockdbInstance.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Times(1).
		DoAndReturn(func(_ any, createUserParams database.CreateUserParams) (database.User, error) {
			mockDBUser := database.User(createUserParams)
			return mockDBUser, nil
		})

	writer := makeRequest("POST", "/user/signup", newUser, map[string]string{}, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusCreated, writer.Code)

}

func TestLoginwithValidCredentials(t *testing.T) {
	loginCredentials := models.LoginData{
		Email:    utils.GenerateRandomEmail(),
		Password: utils.GenerateRandomPassword(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromEmail(gomock.Any(), gomock.Eq(loginCredentials.Email)).
		Times(1).
		DoAndReturn(func(_ any, email string) (database.User, error) {
			mockUser := utils.GenerateMockDatabaseUser()
			salt := utils.GenerateSalt()

			mockUser.Salt = salt
			mockUser.Email = email
			mockUser.HashedPassword = utils.HashPassword(loginCredentials.Password, salt)
			return mockUser, nil
		})

	writer := makeRequest("POST", "/user/login", loginCredentials, map[string]string{}, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]
	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestLoginwithInvalidCredentials(t *testing.T) {
	loginCredentials := models.LoginData{
		Email:    utils.GenerateRandomEmail(),
		Password: utils.GenerateRandomPassword(),
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromEmail(gomock.Any(), gomock.Eq(loginCredentials.Email)).
		Times(1).
		DoAndReturn(func(_ any, email string) (database.User, error) {
			mockUser := utils.GenerateMockDatabaseUser()
			mockUser.Email = loginCredentials.Email
			return mockUser, nil
		})

	writer := makeRequest("POST", "/user/login", loginCredentials, map[string]string{}, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]
	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestListUserSuccess(t *testing.T) {

	mockUser := utils.GenerateMockDatabaseUser()
	mockUser.IsAdminUser = true
	authToken := mockUser.AuthToken
	headers := map[string]string{"AuthToken": authToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromAuthToken(gomock.Any(),
			gomock.Eq(authToken)).
		Times(1).
		Return(mockUser, nil)

	MockdbInstance.EXPECT().
		GetAllUsers(gomock.Any()).
		Times(1).
		Return([]database.GetAllUsersRow{}, nil)

	writer := makeRequest("GET", "/admin/listUser", struct{}{}, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestListUserUnauthorised(t *testing.T) {

	mockUser := utils.GenerateMockDatabaseUser()
	authToken := mockUser.AuthToken
	headers := map[string]string{"AuthToken": authToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromAuthToken(gomock.Any(),
			gomock.Eq(authToken)).
		Times(1).
		Return(mockUser, nil)

	writer := makeRequest("GET", "/admin/listUser", struct{}{}, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
