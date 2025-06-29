package store

import (
	"management/model"
	"management/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

type StoreOperations interface {
	NewStore() error
}

func (store *Postgres) NewStore() error {
	dsn := "host=localhost user=postgres password=191491 dbname=Management port=5432 sslmode=disable"
	utils.Log(model.LogLevelInfo, model.Store, model.NewStore, "database is connecting", nil)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log(model.LogLevelError, model.Store, model.NewStore, "database is not connected", err)
		return err
	} else {
		utils.Log(model.LogLevelInfo, model.Store, model.NewStore, "database connected", nil)
		store.DB = db
	}
	return nil
}
