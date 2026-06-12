package users

import (
	"gorm.io/gorm"
)

var (
	UserRepo    *repositoryUsers
	UserService *serviceUsers
	UserHandler *handlerUsers
)

// InitUsers inicializa todas las capas de usuarios
// Se llama desde main.go con la conexión a la base de datos
func InitUsers(db *gorm.DB) {
	// 1. Primero inicializa el Repository (acceso a BD)
	UserRepo = &repositoryUsers{}

	// 2. Luego el Service (lógica de negocio) pasándole el repo y la bd
	UserService = &serviceUsers{
		repo: UserRepo,
		db:   db,
	}

	// 3. Finalmente el Handler (respuestas HTTP) pasándole el service
	UserHandler = &handlerUsers{
		service: UserService,
	}
}
