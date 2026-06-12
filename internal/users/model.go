package users

import (
	"banc-api/internal/rol"
)

// User es la entidad principal de usuarios en el sistema.
// Contiene los datos esenciales y el rol que controla permisos.
type User struct {
	ID              uint     `gorm:"primaryKey" json:"id"`
	Username        string   `gorm:"unique;not null" json:"username"`
	Email           string   `gorm:"unique;not null" json:"email"`
	Password        string   `gorm:"not null" json:"-"`
	Role            rol.Role `gorm:"not null" json:"role"`
	FechadeCreacion string   `gorm:"not null" json:"fecha_creacion"`
}

// HasPermission comprueba si el usuario tiene el permiso solicitado.
// Esto permite verificar roles antes de ejecutar operaciones sensibles.
func (u *User) HasPermission(permission rol.Permission) bool {
	return u.Role.Can(permission)
}
