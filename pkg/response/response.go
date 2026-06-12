package response

import "github.com/gin-gonic/gin"

// Response representa una respuesta HTTP estándar.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success responde con una respuesta exitosa.
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error responde con una respuesta de error.
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error:   message,
	})
}

// Created responde con 201 Created.
func Created(c *gin.Context, message string, data interface{}) {
	Success(c, 201, message, data)
}

// OK responde con 200 OK.
func OK(c *gin.Context, data interface{}) {
	Success(c, 200, "OK", data)
}

// BadRequest responde con 400 Bad Request.
func BadRequest(c *gin.Context, message string) {
	Error(c, 400, message)
}

// NotFound responde con 404 Not Found.
func NotFound(c *gin.Context, message string) {
	Error(c, 404, message)
}

// InternalServerError responde con 500 Internal Server Error.
func InternalServerError(c *gin.Context, message string) {
	Error(c, 500, message)
}
