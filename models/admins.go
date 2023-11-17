package models

import "github.com/nerd500/axios-cp-wing/internal/database"

type AdminList struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func DBAdminListRowtoAdminListRow(obj database.GetAllAdminUsersRow) AdminList {
	return AdminList{
		Email:     obj.Email,
		FirstName: obj.FirstName,
		LastName:  obj.LastName,
	}
}

func DBAdminListtoAdminList(list []database.GetAllAdminUsersRow) []AdminList {
	var res []AdminList
	for _, obj := range list {
		res = append(res, DBAdminListRowtoAdminListRow(obj))
	}
	return res
}
