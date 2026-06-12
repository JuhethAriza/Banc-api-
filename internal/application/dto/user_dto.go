package dto

import "banc-api/internal/domain/valueobject"

// CreateUserRequest representa la petición para crear un usuario.
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

// UpdateUserRequest representa la petición para actualizar un usuario.
type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// UserResponse representa la respuesta de usuario sin datos sensibles.
type UserResponse struct {
	ID              uint               `json:"id"`
	Username        string             `json:"username"`
	Email           string             `json:"email"`
	Role            valueobject.Role   `json:"role"`
	FechadeCreacion string             `json:"fecha_creacion"`
}
