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

func TestCreateTaskSuccess(t *testing.T) {
	taskCreationRequestBody := models.TaskCreationRequestData{
		Email:    utils.GenerateRandomEmail(),
		Title:    utils.GenerateRandomName(),
		Link:     utils.GenerateRandomLink(),
		Platform: utils.GenerateRandomUUID(),
	}

	authToken := utils.GenerateAuthToken(constants.AuthTokenSize)
	headers := map[string]string{"AuthToken": authToken}

	salt := utils.GenerateSalt()
	mockUser := database.User{
		ID:             uuid.New(),
		Email:          taskCreationRequestBody.Email,
		FirstName:      utils.GenerateRandomName(),
		LastName:       utils.GenerateRandomName(),
		Salt:           salt,
		HashedPassword: utils.HashPassword(utils.GenerateRandomPassword(), salt),
		AuthToken:      authToken,
		IsAdminUser:    true,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromAuthToken(gomock.Any(),
			gomock.Eq(authToken)).
		Times(1).
		DoAndReturn(func(_ any, _ any) (database.User, error) {
			return mockUser, nil
		})

	MockdbInstance.EXPECT().
		CreateTask(gomock.Any(), gomock.Any()).
		Times(1).
		DoAndReturn(func(_ any, createTaskParams database.CreateTaskParams) (database.Task, error) {
			return database.Task(createTaskParams), nil
		})

	writer := makeRequest("POST", "/admin/createTask", taskCreationRequestBody, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestCreateTaskUnAuthorised(t *testing.T) {
	taskCreationRequestBody := models.TaskCreationRequestData{
		Email:    utils.GenerateRandomEmail(),
		Title:    utils.GenerateRandomName(),
		Link:     utils.GenerateRandomLink(),
		Platform: utils.GenerateRandomUUID(),
	}

	authToken := utils.GenerateAuthToken(constants.AuthTokenSize)
	headers := map[string]string{"AuthToken": authToken}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockdbInstance := mockdb.NewMockQuerier(ctrl)

	MockdbInstance.EXPECT().
		GetUserFromAuthToken(gomock.Any(),
			gomock.Eq(authToken)).
		Times(1).
		DoAndReturn(func(_ any, _ any) (database.User, error) {
			return database.User{}, errors.New("Not Found")
		})

	writer := makeRequest("POST", "/admin/createTask", taskCreationRequestBody, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusUnauthorized, writer.Code)
}

func TestGetAllTasksSuccess(t *testing.T) {

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
		DoAndReturn(func(_ any, _ any) (database.User, error) {
			return mockUser, nil
		})

	MockdbInstance.EXPECT().
		GetAllTasks(gomock.Any()).
		Times(1).
		Return([]database.Task{}, nil)

	writer := makeRequest("GET", "/admin/tasks", struct{}{}, headers, MockdbInstance)

	var response map[string]string
	json.Unmarshal(writer.Body.Bytes(), &response)

	error_message, exists := response["error"]

	if exists {
		log.Print(error_message)
	}

	assert.Equal(t, http.StatusOK, writer.Code)
}
