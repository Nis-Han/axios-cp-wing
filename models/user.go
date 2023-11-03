package models

import (
	"github.com/google/uuid"
	"github.com/nerd500/axios-cp-wing/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	AuthToken string    `json:"auth"`
	Password  string    `json:"password"`
}

type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthData struct {
	Email     string `json:"email" binding:"required"`
	AuthToken string `json:"auth"`
}

func DbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		AuthToken: dbUser.AuthToken,
	}
}
