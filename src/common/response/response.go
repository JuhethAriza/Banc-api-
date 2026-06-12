package response

import (
	"net/http"

	"banc-api/src/common/dto"

	"github.com/gin-gonic/gin"
)

// Success responde con éxito
func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, dto.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created responde con recurso creado
func Created(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, dto.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error responde con error
func Error(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, dto.ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// NotFound responde con recurso no encontrado
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, "")
}

// BadRequest responde con error de solicitud
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message, "")
}

// InternalError responde con error interno
func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message, "")
}
