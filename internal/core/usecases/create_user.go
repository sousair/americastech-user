package usecases

import "github.com/sousair/americastech-user/internal/core/entities"

type (
	CreateUserParams struct {
		Name        string
		Email       string
		Password    string
		PhoneNumber string
	}

	CreateUserUseCase interface {
		Create(params CreateUserParams) (*entities.User, error)
	}
)
