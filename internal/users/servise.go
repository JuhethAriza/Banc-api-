package users

import (
	"gorm.io/gorm"
)

type serviceUsers struct {
	repo *repositoryUsers
	db   *gorm.DB
}

// GetUsers obtiene todos los usuarios a través del repository
func (s *serviceUsers) GetUsers() ([]User, error) {
	return s.repo.GetUsers(s.db)
}

// GetUserByID obtiene un usuario específico por su ID
func (s *serviceUsers) GetUserByID(userID uint) (*User, error) {
	return s.repo.GetUserByID(s.db, userID)
}

// CreateUser crea un usuario nuevo
func (s *serviceUsers) CreateUser(user *User) error {
	return s.repo.CreateUser(s.db, user)
}

// UpdateUser actualiza un usuario existente por su ID
func (s *serviceUsers) UpdateUser(userID uint, updatedData *User) (*User, error) {
	return s.repo.UpdateUser(s.db, userID, updatedData)
}

// DeleteUser elimina un usuario por su ID
func (s *serviceUsers) DeleteUser(userID uint) error {
	return s.repo.DeleteUser(s.db, userID)
}
