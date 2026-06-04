package migrations

import (
	"log"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var db *gorm.DB

func MigarateDatabase() {
	err := AutoMigrate()

	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}
}
