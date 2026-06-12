package mapper

import (
	"banc-api/src/infrastructure/db/models"
	"banc-api/src/modules/user/application/dto/response"
	"banc-api/src/modules/user/domain/entities"
)

// ToEntity convierte un modelo de base de datos a entidad de dominio
func ToEntity(model *models.User) *entities.User {
	if model == nil {
		return nil
	}
	return &entities.User{
		ID:              model.ID,
		Username:        model.Username,
		Email:           model.Email,
		Password:        model.Password,
		FechadeCreacion: model.FechadeCreacion,
	}
}

// ToModel convierte una entidad de dominio a modelo de base de datos
func ToModel(entity *entities.User) *models.User {
	if entity == nil {
		return nil
	}
	return &models.User{
		ID:              entity.ID,
		Username:        entity.Username,
		Email:           entity.Email,
		Password:        entity.Password,
		FechadeCreacion: entity.FechadeCreacion,
	}
}

// ToResponse convierte una entidad a respuesta DTO
func ToResponse(entity *entities.User) *response.UserResponse {
	if entity == nil {
		return nil
	}
	return &response.UserResponse{
		ID:              entity.ID,
		Username:        entity.Username,
		Email:           entity.Email,
		FechadeCreacion: entity.FechadeCreacion.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// ToResponseList convierte una lista de entidades a lista de respuestas
func ToResponseList(entities []entities.User) []response.UserResponse {
	responses := make([]response.UserResponse, len(entities))
	for i, entity := range entities {
		responses[i] = *ToResponse(&entity)
	}
	return responses
}
