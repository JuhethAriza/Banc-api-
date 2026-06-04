package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(GetDatabaseURL func() string) error {

	db, err := gorm.Open(
		postgres.Open(GetDatabaseURL()),
		&gorm.Config{},
	)

	if err != nil {
		return err
	}

	DB = db

	return nil
}
