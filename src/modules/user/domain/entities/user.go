package entities

import "time"

// User es la entidad principal de usuarios en el sistema
type User struct {
	ID              uint      `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Password        string    `json:"-"`
	FechadeCreacion time.Time `json:"fecha_creacion"`
}
