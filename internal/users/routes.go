package users

import (
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registra todas las rutas de usuarios
// Se llama desde internal/init.go
func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", UserHandler.GetUsers)
	r.GET("/users/:id", UserHandler.GetUserByID)
	r.POST("/users", UserHandler.CreateUser)
	r.PUT("/users/:id", UserHandler.UpdateUser)
	r.DELETE("/users/:id", UserHandler.DeleteUser)
}
