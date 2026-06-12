package dto

import "banc-api/internal/domain/valueobject"

// CreateAccountRequest representa la petición para crear una cuenta.
type CreateAccountRequest struct {
	UserID      uint   `json:"user_id" binding:"required"`
	Token       string `json:"token" binding:"required"`
	TypeAccount string `json:"tipo_cuenta" binding:"required"`
}

// UpdateAccountRequest representa la petición para actualizar una cuenta.
type UpdateAccountRequest struct {
	TipoCuenta   string  `json:"tipo_cuenta"`
	Saldo        float64 `json:"saldo"`
	Estado       string  `json:"estado"`
	NumeroCuenta string  `json:"numero_cuenta"`
}

// AccountResponse representa la respuesta de cuenta.
type AccountResponse struct {
	ID           uint                    `json:"id"`
	UserID       uint                    `json:"user_id"`
	TipoCuenta   valueobject.TypeAccount `json:"tipo_cuenta"`
	Saldo        float64                 `json:"saldo"`
	Estado       string                  `json:"estado"`
	NumeroCuenta string                  `json:"numero_cuenta"`
}
