package entity

// User es la entidad principal de usuarios en el sistema.
// Contiene los datos esenciales y el rol que controla permisos.
type User struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Username        string `gorm:"unique;not null" json:"username"`
	Email           string `gorm:"unique;not null" json:"email"`
	Password        string `gorm:"not null" json:"-"`
	FechadeCreacion string `gorm:"not null" json:"fecha_creacion"`
}
