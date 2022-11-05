package repository

import (
	"dibagi/config"
	"dibagi/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DB_CONFIG), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to Repository", err)
		return nil, err
	}

	err = db.Debug().AutoMigrate(models.User{})
	if err != nil {
		fmt.Println("error migrating Repository", err)
		return nil, err
	}
	return db, nil
}
