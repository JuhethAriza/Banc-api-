package account

import "gorm.io/gorm"

type AccountRepository struct{}

func (r *AccountRepository) PostAccount(db *gorm.DB, account *Account) error {
	return db.Create(account).Error
}

func (r *AccountRepository) PutchAccount(db *gorm.DB, account *Account) error {
	return db.Save(account).Error
}
