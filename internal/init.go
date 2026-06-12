package internal

import (
	"banc-api/internal/application/usecase"
	"banc-api/internal/domain/repository"
	"banc-api/internal/interface/http"
	"banc-api/internal/infrastructure/persistence"
	"banc-api/pkg/database"

	"github.com/gin-gonic/gin"
)

// Container de dependencias de la aplicación.
type AppContainer struct {
	UserRepo       repository.UserRepository
	AccountRepo    repository.AccountRepository
	UserUseCase    *usecase.UserUseCase
	AccountUseCase *usecase.AccountUseCase
	UserHandler    *http.UserHandler
	AccountHandler *http.AccountHandler
}

// NewAppContainer inicializa todas las dependencias de la aplicación.
// Sigue el principio de inversión de dependencias: las capas superiores
// dependen de abstracciones, no de implementaciones.
func NewAppContainer(db *database.DB) *AppContainer {
	// Capa de Infraestructura: implementaciones concretas
	userRepo := persistence.NewUserRepository(db.Instance())
	accountRepo := persistence.NewAccountRepository(db.Instance())

	// Capa de Aplicación: casos de uso
	userUseCase := usecase.NewUserUseCase(userRepo)
	accountUseCase := usecase.NewAccountUseCase(accountRepo)

	// Capa de Interfaz: handlers HTTP
	userHandler := http.NewUserHandler(userUseCase)
	accountHandler := http.NewAccountHandler(accountUseCase)

	return &AppContainer{
		UserRepo:       userRepo,
		AccountRepo:    accountRepo,
		UserUseCase:    userUseCase,
		AccountUseCase: accountUseCase,
		UserHandler:    userHandler,
		AccountHandler: accountHandler,
	}
}

// SetupRoutes configura todas las rutas de la aplicación.
func SetupRoutes(container *AppContainer) *gin.Engine {
	router := gin.Default()

	// Registrar rutas
	http.RegisterUserRoutes(router, container.UserHandler)
	http.RegisterAccountRoutes(router, container.AccountHandler)

	return router
}
