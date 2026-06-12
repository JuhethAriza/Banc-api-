package account

import (
	"banc-api/internal/typeaccount"
	"errors"

	"gorm.io/gorm"
)

type CreateAccountRequest struct {
	UserID      uint
	Token       string
	TypeAccount string
}

type ServiceAccount struct {
	repository AccountRepository
}

func (s *ServiceAccount) PostAccount(db *gorm.DB, account *Account, req CreateAccountRequest) error {
	// basic validation: TypeAccount must be non-empty
	if req.TypeAccount == "" {
		return errors.New("tipo de cuenta inválido")
	}
	newAccount := Account{
		UserID:     req.UserID,
		Token:      req.Token,
		TipoCuenta: typeaccount.TypeAccount(req.TypeAccount),
	}
	return s.repository.PostAccount(db, &newAccount)
}

func (s *ServiceAccount) PutchAccount(db *gorm.DB, account *Account) error {
	return s.repository.PutchAccount(db, account)
}
