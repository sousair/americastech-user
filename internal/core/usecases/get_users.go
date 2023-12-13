package usecases

import "github.com/sousair/americastech-user/internal/core/entities"

type (
	GetUsersUseCase interface {
		Get() ([]*entities.User, error)
	}
)
