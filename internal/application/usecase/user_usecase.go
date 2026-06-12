package usecase

import (
	"banc-api/internal/application/dto"
	"banc-api/internal/domain/entity"
	"banc-api/internal/domain/repository"
	"banc-api/internal/domain/valueobject"
	"errors"
	"time"
)

// UserRepository es la interfaz de repositorio que usamos (inversión de dependencias).
type UserRepository interface {
	repository.UserRepository
}

// UserUseCase implementa los casos de uso de usuarios.
type UserUseCase struct {
	userRepo UserRepository
}

// NewUserUseCase crea un nuevo caso de uso de usuarios.
func NewUserUseCase(userRepo UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// GetUsers obtiene todos los usuarios.
func (uc *UserUseCase) GetUsers() ([]dto.UserResponse, error) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.UserResponse, len(users))
	for i, u := range users {
		responses[i] = dto.UserResponse{
			ID:              u.ID,
			Username:        u.Username,
			Email:           u.Email,
			Role:            u.Role,
			FechadeCreacion: u.FechadeCreacion,
		}
	}
	return responses, nil
}

// GetUserByID obtiene un usuario por su ID.
func (uc *UserUseCase) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Role:            user.Role,
		FechadeCreacion: user.FechadeCreacion,
	}
	return &response, nil
}

// CreateUser crea un nuevo usuario.
func (uc *UserUseCase) CreateUser(req dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Validar rol
	newRole, validRole := valueobject.ParseRole(req.Role)
	if !validRole && req.Role != "" {
		return nil, errors.New("rol de usuario inválido")
	}
	if req.Role == "" {
		newRole = valueobject.RoleUser
	}

	// Crear entidad
	user := &entity.User{
		Username:        req.Username,
		Email:           req.Email,
		Password:        req.Password, // TODO: hashear password
		Role:            newRole,
		FechadeCreacion: time.Now().Format(time.RFC3339),
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Role:            user.Role,
		FechadeCreacion: user.FechadeCreacion,
	}
	return &response, nil
}

// UpdateUser actualiza un usuario existente.
func (uc *UserUseCase) UpdateUser(id uint, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Actualizar campos
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password // TODO: hashear password
	}
	if req.Role != "" {
		updatedRole, validRole := valueobject.ParseRole(req.Role)
		if !validRole || updatedRole == valueobject.RoleGuest {
			return nil, errors.New("rol de usuario inválido")
		}
		user.Role = updatedRole
	}

	if err := uc.userRepo.Update(user); err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Role:            user.Role,
		FechadeCreacion: user.FechadeCreacion,
	}
	return &response, nil
}

// DeleteUser elimina un usuario.
func (uc *UserUseCase) DeleteUser(id uint) error {
	return uc.userRepo.Delete(id)
}
