package migrations

import (
	"banc-api/internal/users"
	"banc-api/pkg/database"
)

func RunMigrations() error {

	err := database.DB.AutoMigrate(
		&users.User{},
	)

	if err != nil {
		return err
	}

	return nil
}
