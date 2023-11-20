package handlers_test

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/nerd500/axios-cp-wing/internal/database"
	mockdb "github.com/nerd500/axios-cp-wing/internal/database/mock"
	"github.com/nerd500/axios-cp-wing/models"
	"github.com/nerd500/axios-cp-wing/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestListAdminSuccess(t *testing.T) {

	mockUser := utils.GenerateMockDatabaseUser()
	mockUser.IsAdminUser = true
	mockUser.Email = os.Getenv("ROOT_USER_EMAIL")
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
		GetAllAdminUsers(gomock.Any()).
		Times(1).
		Return([]database.GetAllAdminUsersRow{}, nil)

	writer := makeRequest("GET", "/root/listAdmin", struct{}{}, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestListAdminUnauthorised(t *testing.T) {

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

	writer := makeRequest("GET", "/root/listAdmin", struct{}{}, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestEditAdminPermissionsSuccess(t *testing.T) {
	mockUser := utils.GenerateMockDatabaseUser()
	mockUser.IsAdminUser = true
	mockUser.Email = os.Getenv("ROOT_USER_EMAIL")
	authToken := mockUser.AuthToken
	headers := map[string]string{"AuthToken": authToken}

	params := models.EditAdminAccessParams{}
	params.Email = utils.GenerateRandomEmail()
	params.IsAdminUser = true

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromAuthToken(gomock.Any(),
			gomock.Eq(authToken)).
		Times(2).
		Return(mockUser, nil)

	MockdbInstance.EXPECT().
		GetUserFromEmail(gomock.Any(), params.Email).
		Times(2).
		Return(utils.GenerateMockDatabaseUser(), nil)

	MockdbInstance.EXPECT().
		EditAdminAccess(gomock.Any(), gomock.Any()).
		Times(2).
		DoAndReturn(func(_ any, _ any) (database.User, error) {
			mockReturnUser := utils.GenerateMockDatabaseUser()
			mockReturnUser.IsAdminUser = params.IsAdminUser
			mockReturnUser.Email = params.Email
			return mockReturnUser, nil
		})
	writer := makeRequest("PATCH", "/root/updateAdminPermission", params, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}
	assert.Equal(t, http.StatusOK, writer.Code)

	params.IsAdminUser = false

	writer = makeRequest("PATCH", "/root/updateAdminPermission", params, headers, MockdbInstance)

	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists = response["error"]

	if exists {
		log.Print(error_message)
	}
	assert.Equal(t, http.StatusOK, writer.Code)

}
