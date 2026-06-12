package migrations

import (
	"banc-api/internal/domain/entity"

	"gorm.io/gorm"
)

// RunMigrations ejecuta las migraciones de la base de datos.
func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
	)

	if err != nil {
		return err
	}

	return nil
}
