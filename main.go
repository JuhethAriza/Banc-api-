package main

import (
	"banc-api/config"
	"banc-api/internal"
	"banc-api/migrations"
	"banc-api/pkg/database"
	"log"
)

func main() {
	// 1. Cargar variables de entorno
	config.LoadEnv()

	// 2. Conectar a la base de datos
	db, err := database.ConnectDatabase(config.GetDatabaseURL)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos: " + err.Error())
	}

	// 3. Ejecutar migraciones
	if err := migrations.RunMigrations(db.Instance()); err != nil {
		log.Fatal("No se pudieron ejecutar las migraciones: " + err.Error())
	}

	// 4. Inicializar contenedor de dependencias
	container := internal.NewAppContainer(db)

	// 5. Configurar rutas (registrando handlers desde el container)
	r := internal.SetupRoutes(container)

	// 6. Iniciar servidor
	log.Println("Servidor iniciado en :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: " + err.Error())
	}
}
