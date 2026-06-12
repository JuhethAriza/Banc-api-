package http

import (
	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registra las rutas de usuarios en el router.
func RegisterUserRoutes(r *gin.Engine, handler *UserHandler) {
	users := r.Group("/users")
	{
		users.GET("", handler.GetUsers)
		users.GET("/:id", handler.GetUserByID)
		users.POST("", handler.CreateUser)
		users.PUT("/:id", handler.UpdateUser)
		users.DELETE("/:id", handler.DeleteUser)
	}
}

// RegisterAccountRoutes registra las rutas de cuentas en el router.
func RegisterAccountRoutes(r *gin.Engine, handler *AccountHandler) {
	accounts := r.Group("/accounts")
	{
		accounts.GET("", handler.GetAccounts)
		accounts.GET("/:id", handler.GetAccountByID)
		accounts.POST("", handler.CreateAccount)
		accounts.PUT("/:id", handler.UpdateAccount)
		accounts.DELETE("/:id", handler.DeleteAccount)
	}
}
