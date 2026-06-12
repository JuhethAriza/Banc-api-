package persistence

import (
	"banc-api/internal/domain/entity"
	"banc-api/internal/domain/repository"
	"gorm.io/gorm"
)

// userRepositoryImpl implementa UserRepository usando GORM.
type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva implementación de UserRepository.
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepositoryImpl) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepositoryImpl) Delete(id uint) error {
	res := r.db.Delete(&entity.User{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
