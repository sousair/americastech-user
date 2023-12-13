package usecases

import "github.com/sousair/americastech-user/internal/core/entities"

type (
	UpdateUserParams struct {
		ID          string
		Email       string
		Name        string
		PhoneNumber string
	}

	UpdateUserUseCase interface {
		Update(params UpdateUserParams) (*entities.User, error)
	}
)
