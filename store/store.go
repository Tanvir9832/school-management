package store

import (
	"fmt"
	"log"

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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error in the db connection : %v", err)
		return err
	} else {
		fmt.Println(db)
		store.DB = db
	}
	return nil
}
