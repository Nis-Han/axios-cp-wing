package models

import "github.com/nerd500/axios-cp-wing/internal/database"

type EditAdminAccessParams struct {
	IsAdminUser bool   `json:"is_admin_user"`
	Email       string `json:"email"`
}

func EditAdminAccessParamsToDBModel(obj EditAdminAccessParams) database.EditAdminAccessParams {
	return database.EditAdminAccessParams{
		IsAdminUser: obj.IsAdminUser,
		Email:       obj.Email,
	}
}
