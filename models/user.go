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

func DbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		AuthToken: dbUser.AuthToken,
	}
}

type UserList struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func DBAdminListRowtoAdminListRow(obj database.GetAllAdminUsersRow) UserList {
	return UserList{
		Email:     obj.Email,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
	}
}

func DBAdminListtoAdminList(list []database.GetAllAdminUsersRow) []UserList {
	var res []UserList
	for _, obj := range list {
		res = append(res, DBAdminListRowtoAdminListRow(obj))
	}
	return res
}

func DBAUserListRowtoUserListRow(obj database.GetAllUsersRow) UserList {
	return UserList{
		Email:     obj.Email,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
	}
}

func DBUserListtoUserList(list []database.GetAllUsersRow) []UserList {
	var res []UserList
	for _, obj := range list {
		res = append(res, DBAUserListRowtoUserListRow(obj))
	}
	return res
}
