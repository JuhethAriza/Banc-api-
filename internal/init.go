package internal

import (
	"banc-api/internal/users"
	"banc-api/pkg/database"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

// SetupRoutes inicializa todas las capas y registra las rutas
// Se llama desde main.go
func SetupRoutes() *gin.Engine {
	router = gin.Default()

	// Inicializar los usuarios con la conexión a la BD
	users.InitUsers(database.DB)

	// Registrar rutas de usuarios
	users.RegisterUserRoutes(router)

	// Aquí irían otras rutas (auth, accounts, transfers, etc)

	return router
}
