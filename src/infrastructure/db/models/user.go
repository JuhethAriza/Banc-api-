package models

import "time"

// User es el modelo de base de datos para la tabla users
type User struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Username        string    `gorm:"unique;not null" json:"username"`
	Email           string    `gorm:"unique;not null" json:"email"`
	Password        string    `gorm:"not null" json:"-"`
	FechadeCreacion time.Time `gorm:"not null" json:"fecha_creacion"`
}

// TableName especifica el nombre de la tabla
func (User) TableName() string {
	return "users"
}
