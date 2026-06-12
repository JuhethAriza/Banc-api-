package users

import (
	"net/http"
	"strconv"
	"time"

	"banc-api/internal/rol"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetUsers() ([]User, error)
	GetUserByID(userID uint) (*User, error)
	CreateUser(user *User) error
	UpdateUser(userID uint, updatedData *User) (*User, error)
	DeleteUser(userID uint) error
}

type handlerUsers struct {
	service UserServiceInterface
}

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

type updateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// GetUsers maneja la petición HTTP GET /users
func (h *handlerUsers) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID maneja la petición HTTP GET /users/:id
func (h *handlerUsers) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario requerido"})
		return
	}

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuario"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser maneja la petición HTTP POST /users
func (h *handlerUsers) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	newRole, validRole := rol.ParseRole(req.Role)
	if !validRole && req.Role != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rol de usuario inválido"})
		return
	}

	if req.Role == "" {
		newRole = rol.RoleUser
	}

	user := &User{
		Username:        req.Username,
		Email:           req.Email,
		Password:        req.Password,
		Role:            newRole,
		FechadeCreacion: time.Now().Format(time.RFC3339),
	}

	if err := h.service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser maneja la petición HTTP PUT /users/:id
func (h *handlerUsers) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario requerido"})
		return
	}

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	updatedData := &User{}
	if req.Username != "" {
		updatedData.Username = req.Username
	}
	if req.Email != "" {
		updatedData.Email = req.Email
	}
	if req.Password != "" {
		updatedData.Password = req.Password
	}
	if req.Role != "" {
		updatedRole, validRole := rol.ParseRole(req.Role)
		if !validRole || updatedRole == rol.RoleGuest {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Rol de usuario inválido"})
			return
		}
		updatedData.Role = updatedRole
	}

	user, err := h.service.UpdateUser(uint(id), updatedData)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser maneja la petición HTTP DELETE /users/:id
func (h *handlerUsers) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario requerido"})
		return
	}

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	if err := h.service.DeleteUser(uint(id)); err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
		return
	}

	c.Status(http.StatusNoContent)
}
