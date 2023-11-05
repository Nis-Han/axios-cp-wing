package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

type TaskCreationRequestData struct {
	Email    string   `json:"email" binding:"required"`
	Title    string   `json:"title"`
	Link     string   `json:"link"`
	Tags     []string `json:"tags"`
	Platform string   `json:"platform"`
}

type Task struct {
	ID           uuid.UUID `json:"id"`
	CreatedBy    string    `json:"created_by"`
	CreatedAt    time.Time `json:"created_at"`
	LastEditedBy string    `json:"last_edited_by"`
	LastEditedAt time.Time `json:"last_edited_at"`
	Title        string    `json:"title"`
	Link         string    `json:"link"`
	Tags         []string  `json:"tags"`
	Platform     string    `json:"platform"`
}

func DbTaskToTask(dbTask database.Task) Task {
	return Task{
		ID:           dbTask.ID,
		CreatedBy:    dbTask.CreatedBy,
		CreatedAt:    dbTask.CreatedAt,
		LastEditedBy: dbTask.LastEditedBy,
		LastEditedAt: dbTask.LastEditedAt,
		Title:        dbTask.Title,
		Link:         dbTask.Link,
		Tags:         dbTask.Tags,
		Platform:     dbTask.Platform,
	}
}
