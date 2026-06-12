package users

type UserDto struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	FechadeCreacion string `json:"fecha_creacion"`
}

// ToDTO devuelve una representación segura del usuario para devolver en la API.
// Oculta campos sensibles como la contraseña.
func (u *User) ToDTO() UserDto {
	return UserDto{
		ID:              u.ID,
		Username:        u.Username,
		Email:           u.Email,
		Role:            u.Role.String(),
		FechadeCreacion: u.FechadeCreacion,
	}
}
