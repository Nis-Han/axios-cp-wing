package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/models"
)

func (api *Api) CreateTask(c *gin.Context) {

	decoder := json.NewDecoder(c.Request.Body)

	taskCreationRequestData := models.TaskCreationRequestData{}

	err := decoder.Decode(&taskCreationRequestData)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Ill-formatted request body")
		return
	}
	var newTaskData database.CreateTaskParams

	userData := c.MustGet("userData")
	user, ok := userData.(database.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	newTaskData.ID = uuid.New()
	newTaskData.CreatedAt = time.Now()
	newTaskData.LastEditedAt = time.Now()
	newTaskData.CreatedBy = user.ID
	newTaskData.LastEditedBy = user.ID
	newTaskData.Title = taskCreationRequestData.Title
	newTaskData.Link = taskCreationRequestData.Link
	newTaskData.Platform = taskCreationRequestData.Platform

	dbTask, err := api.DB.CreateTask(c.Request.Context(), newTaskData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt create task"})
		return
	}

	c.JSON(http.StatusCreated, models.DbTaskToTask(dbTask))
}

func (api *Api) GetAllTasks(c *gin.Context) {

	tasks, err := api.DB.GetAllTasks(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt get task"})
		return
	}

	c.JSON(http.StatusOK, models.DbTaskListToTaskList(tasks))
}
