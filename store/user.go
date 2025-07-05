package store

import (
	"management/model"
	"management/utils"
)

func (store Postgres) CreateUser(user *model.User) error {
	utils.Log(model.LogLevelInfo, model.Store, model.CreateUser, "Creating new user", nil)
	response := store.DB.Create(user)
	if response.Error != nil {
		utils.Log(model.LogLevelError, model.Store, model.CreateUser, "error while creating new user", response.Error)
		return response.Error
	}
	utils.Log(model.LogLevelError, model.Store, model.CreateUser, "Created new user", nil)
	return nil
}
func (store Postgres) UpdateUser(user *model.User) error {
	return nil
}
