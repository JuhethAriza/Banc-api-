package repositories

import "banc-api/src/modules/user/domain/entities"

type UserRepository interface {
	GetAll() ([]entities.User, error)
}
