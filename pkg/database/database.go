package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es un wrapper alrededor de gorm.DB para facilitar testing y abstracción.
type DB struct {
	instance *gorm.DB
}

// Instance devuelve la instancia subyacente de gorm.DB.
func (d *DB) Instance() *gorm.DB {
	return d.instance
}

// ConnectDatabase conecta a la base de datos y devuelve un wrapper DB.
func ConnectDatabase(getDatabaseURL func() string) (*DB, error) {
	db, err := gorm.Open(
		postgres.Open(getDatabaseURL()),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return &DB{instance: db}, nil
}
