package config

import "os"

func GetDatabaseURL() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}
	return "postgresql://postgres:postgres@localhost:5432/neondb?sslmode=disable"
}
