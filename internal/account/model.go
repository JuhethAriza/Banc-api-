package account

import (
	"banc-api/internal/typeaccount"
	"time"
)

type Account struct {
	ID           uint                    `gorm:"primaryKey" json:"id"`
	UserID       uint                    `json:"user_id"`
	NumeroCuenta string                  `json:"numero_cuenta" gorm:"unique;not null"`
	TipoCuenta   typeaccount.TypeAccount `json:"tipo_cuenta"`
	Saldo        float64                 `json:"saldo"`
	Estado       string                  `json:"estado"`
	// Otros campos relevantes del modelo de cuenta
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
	Token              string    `json:"token" gorm:"unique;not null"`
}

type StateAccount struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Activo    bool `json:"activo"`
	Bloqueado bool `json:"bloqueado"`
	Cerrado   bool `json:"cerrado"`
}

func (a *Account) permissions(t typeaccount.TypeAccount) bool {
	// Simple permission check: compare provided type with the account's type
	return t == typeaccount.TypeAccount(a.TipoCuenta)
}
