package repository

import (
	"banc-api/internal/domain/entity"
)

// UserRepository define el contrato para el acceso a datos de usuarios.
// Esta interfaz vive en el dominio para invertir dependencias.
type UserRepository interface {
	// GetAll obtiene todos los usuarios
	GetAll() ([]entity.User, error)
	// GetByID obtiene un usuario por su ID
	GetByID(id uint) (*entity.User, error)
	// GetByEmail obtiene un usuario por su email
	GetByEmail(email string) (*entity.User, error)
	// Create crea un nuevo usuario
	Create(user *entity.User) error
	// Update actualiza un usuario existente
	Update(user *entity.User) error
	// Delete elimina un usuario por su ID
	Delete(id uint) error
}
