package account

import "banc-api/internal/typeaccount"

type ResponseAccount struct {
	ID         uint                    `json:"id"`
	UserID     uint                    `json:"user_id"`
	TipoCuenta typeaccount.TypeAccount `json:"tipo_cuenta"`
	Saldo      float64                 `json:"saldo"`
	Estado     string                  `json:"estado"`
}
