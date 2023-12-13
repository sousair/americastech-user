package usecases

import "github.com/sousair/americastech-user/internal/core/entities"

type (
	CreateUserParams struct {
		Name        string `json:"username"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		PhoneNumber string `json:"phone_number"`
	}

	CreateUserUseCase interface {
		Create(params CreateUserParams) (*entities.User, error)
	}
)
