package repository

import (
	"banc-api/internal/domain/entity"
)

// AccountRepository define el contrato para el acceso a datos de cuentas.
type AccountRepository interface {
	// GetAll obtiene todas las cuentas
	GetAll() ([]entity.Account, error)
	// GetByID obtiene una cuenta por su ID
	GetByID(id uint) (*entity.Account, error)
	// GetByUserID obtiene todas las cuentas de un usuario
	GetByUserID(userID uint) ([]entity.Account, error)
	// Create crea una nueva cuenta
	Create(account *entity.Account) error
	// Update actualiza una cuenta existente
	Update(account *entity.Account) error
	// Delete elimina una cuenta por su ID
	Delete(id uint) error
}
