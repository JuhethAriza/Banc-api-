package response

// UserResponse representa la respuesta de usuario sin datos sensibles
type UserResponse struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	FechadeCreacion string `json:"fecha_creacion"`
}

// LoginResponse representa la respuesta de inicio de sesión
type LoginResponse struct {
	Token     string       `json:"token"`
	ExpiresIn int64        `json:"expires_in"`
	User      UserResponse `json:"user"`
}
