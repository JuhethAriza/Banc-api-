package users

import (
	"gorm.io/gorm"
)

type repositoryUsers struct{}

// GetUsers obtiene todos los usuarios de la base de datos
// Devuelve: lista de usuarios y error
// El Repository NO responde HTTP, solo obtiene datos
func (r *repositoryUsers) GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID obtiene un usuario específico por su ID
func (r *repositoryUsers) GetUserByID(db *gorm.DB, userID uint) (*User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser crea un nuevo usuario en la base de datos
func (r *repositoryUsers) CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// UpdateUser actualiza un usuario existente por su ID
func (r *repositoryUsers) UpdateUser(db *gorm.DB, userID uint, updatedData *User) (*User, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&user).Updates(updatedData).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser elimina un usuario por su ID
func (r *repositoryUsers) DeleteUser(db *gorm.DB, userID uint) error {
	res := db.Delete(&User{}, userID)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
