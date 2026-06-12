package entity

import (
	"banc-api/internal/domain/valueobject"
	"time"
)

// Account representa una cuenta bancaria en el sistema.
type Account struct {
	ID                 uint                  `gorm:"primaryKey" json:"id"`
	UserID             uint                  `json:"user_id"`
	NumeroCuenta       string                `json:"numero_cuenta" gorm:"unique;not null"`
	TipoCuenta         valueobject.TypeAccount `json:"tipo_cuenta"`
	Saldo              float64               `json:"saldo"`
	Estado             string                `json:"estado"`
	FechaCreacion      time.Time             `json:"fecha_creacion"`
	FechaActualizacion time.Time             `json:"fecha_actualizacion"`
	Token              string                `json:"token" gorm:"unique;not null"`
}

// StateAccount representa el estado de una cuenta.
type StateAccount struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	Activo    bool `json:"activo"`
	Bloqueado bool `json:"bloqueado"`
	Cerrado   bool `json:"cerrado"`
}

// HasPermission verifica permisos basados en el tipo de cuenta.
func (a *Account) HasPermission(t valueobject.TypeAccount) bool {
	return t == a.TipoCuenta
}
