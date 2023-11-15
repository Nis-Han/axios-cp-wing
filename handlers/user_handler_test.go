package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/nerd500/axios-cp-wing/constants"
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
		GetUser(gomock.Any(), gomock.Eq(newUser.Email)).
		Times(1).
		Return(database.User{}, errors.New("User Not Found Error"))

	MockdbInstance.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Times(1).
		DoAndReturn(func(_ any, createUserParams database.CreateUserParams) (database.User, error) {
			mockDBUser := database.User(createUserParams)
			return mockDBUser, nil
		})

	database.DBInstance = MockdbInstance

	writer := makeRequest("POST", "/user/signup", newUser, map[string]string{})

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
		GetUser(gomock.Any(), gomock.Eq(loginCredentials.Email)).
		Times(1).
		DoAndReturn(func(_ any, email string) (database.User, error) {
			salt := utils.GenerateSalt()
			return database.User{
				ID:             uuid.New(),
				Email:          email,
				FirstName:      utils.GenerateRandomName(),
				LastName:       utils.GenerateRandomName(),
				Salt:           salt,
				HashedPassword: utils.HashPassword(loginCredentials.Password, salt),
				AuthToken:      utils.GenerateAuthToken(constants.AuthTokenSize),
				IsAdminUser:    false,
			}, nil
		})

	database.DBInstance = MockdbInstance

	writer := makeRequest("POST", "/user/login", loginCredentials, map[string]string{})

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
		GetUser(gomock.Any(), gomock.Eq(loginCredentials.Email)).
		Times(1).
		DoAndReturn(func(_ any, email string) (database.User, error) {
			salt := utils.GenerateSalt()
			return database.User{
				ID:             uuid.New(),
				Email:          email,
				FirstName:      utils.GenerateRandomName(),
				LastName:       utils.GenerateRandomName(),
				Salt:           salt,
				HashedPassword: utils.HashPassword(utils.GenerateRandomPassword(), salt),
				AuthToken:      utils.GenerateAuthToken(100),
				IsAdminUser:    false,
			}, nil
		})

	database.DBInstance = MockdbInstance

	writer := makeRequest("POST", "/user/login", loginCredentials, map[string]string{})

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]
	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}
