package valueobject

import "strings"

// TypeAccount representa el tipo de cuenta bancaria.
type TypeAccount string

const (
	TypeAccountAhorro    TypeAccount = "Ahorro"
	TypeAccountCorriente TypeAccount = "Corriente"
)

// String convierte el tipo de cuenta a su valor de texto.
func (t TypeAccount) String() string {
	return string(t)
}

// ParseTypeAccount normaliza un valor de texto a un tipo de cuenta válido.
func ParseTypeAccount(value string) (TypeAccount, bool) {
	normalized := TypeAccount(strings.ToLower(strings.TrimSpace(value)))
	switch normalized {
	case "ahorro":
		return TypeAccountAhorro, true
	case "corriente":
		return TypeAccountCorriente, true
	default:
		return TypeAccountCorriente, false
	}
}
