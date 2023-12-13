package usecases

import "github.com/sousair/americastech-user/internal/core/entities"

type (
	GetUserParams struct {
		ID string
	}

	GetUserUseCase interface {
		Get(params GetUserParams) (*entities.User, error)
	}
)
