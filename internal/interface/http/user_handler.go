package http

import (
	"banc-api/internal/application/dto"
	"banc-api/internal/application/usecase"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler maneja las peticiones HTTP de usuarios.
type UserHandler struct {
	useCase *usecase.UserUseCase
}

// NewUserHandler crea un nuevo handler de usuarios.
func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		useCase: uc,
	}
}

// GetUsers maneja GET /users.
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.useCase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUserByID maneja GET /users/:id.
func (h *UserHandler) GetUserByID(c *gin.Context) {
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

	user, err := h.useCase.GetUserByID(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser maneja POST /users.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos", "details": err.Error()})
		return
	}

	user, err := h.useCase.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente", "data": user})
}

// UpdateUser maneja PUT /users/:id.
func (h *UserHandler) UpdateUser(c *gin.Context) {
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

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	user, err := h.useCase.UpdateUser(uint(id), req)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente", "data": user})
}

// DeleteUser maneja DELETE /users/:id.
func (h *UserHandler) DeleteUser(c *gin.Context) {
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

	if err := h.useCase.DeleteUser(uint(id)); errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
		return
	}

	c.Status(http.StatusNoContent)
}
