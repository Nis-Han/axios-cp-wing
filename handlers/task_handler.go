package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nerd500/axios-cp-wing/internal/database"
	"github.com/nerd500/axios-cp-wing/models"
)

func CreateTask(c *gin.Context) {
	var taskCreationRequestData models.TaskCreationRequestData
	err := c.ShouldBindJSON(&taskCreationRequestData)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Ill-formatted request body")
		return
	}

	var newTaskData database.CreateTaskParams

	newTaskData.ID = uuid.New()
	newTaskData.CreatedAt = time.Now()
	newTaskData.LastEditedAt = time.Now()
	newTaskData.CreatedBy = taskCreationRequestData.Email
	newTaskData.LastEditedBy = taskCreationRequestData.Email
	newTaskData.Title = taskCreationRequestData.Title
	newTaskData.Link = taskCreationRequestData.Link
	newTaskData.Tags = taskCreationRequestData.Tags
	newTaskData.Platform = taskCreationRequestData.Platform

	dbTask, err := database.DBInstance.CreateTask(c.Request.Context(), newTaskData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt create task"})
		return
	}

	c.JSON(http.StatusCreated, models.DbTaskToTask(dbTask))
}
