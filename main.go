package main

import (
	"banc-api/config"
	"banc-api/internal"
	"banc-api/migrations"
	"banc-api/pkg/database"
)

func main() {
	// 1. Cargar variables de entorno
	config.LoadEnv()

	// 2. Conectar a la base de datos
	err := database.ConnectDatabase(config.GetDatabaseURL)
	if err != nil {
		panic("No se pudo conectar a la base de datos: " + err.Error())
	}

	// 3. Ejecutar migraciones
	if err := migrations.RunMigrations(); err != nil {
		panic("No se pudieron ejecutar las migraciones: " + err.Error())
	}

	// 4. Configurar rutas (aquí se inicializan todas las capas)
	r := internal.SetupRoutes()

	// 5. Iniciar servidor
	r.Run(":8080")
}
