package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load carga las variables de entorno desde .env
func Load() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found; using environment variables")
	}
}

// GetDatabaseURL retorna la URL de conexión a la base de datos
func GetDatabaseURL() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}
	return "postgresql://postgres:postgres@localhost:5432/neondb?sslmode=disable"
}

// GetServerPort retorna el puerto del servidor
func GetServerPort() string {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		return port
	}
	return "8080"
}
