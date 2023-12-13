package app_usecases

import (
	"errors"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/application/providers/cryptography"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	createUserUseCase struct {
		userRepository repositories.UserRepository
		cryptoProvider cryptography.CryptoProvider
	}
)

func NewCreateUserUseCase(userRepo repositories.UserRepository, cryptoProvider cryptography.CryptoProvider) usecases.CreateUserUseCase {
	return &createUserUseCase{
		userRepository: userRepo,
		cryptoProvider: cryptoProvider,
	}
}

func (uc createUserUseCase) Create(params usecases.CreateUserParams) (*entities.User, error) {
	emailAlreadyExists, _ := uc.userRepository.FindOneBy(map[string]interface{}{
		"email": params.Email,
	})

	if emailAlreadyExists != nil {
		return nil, custom_errors.NewEmailAlreadyExistsError(errors.New(""), params.Email)
	}

	encryptedPassword, err := uc.cryptoProvider.Hash(params.Password)

	if err != nil {
		return nil, err
	}

	user, err := uc.userRepository.Create(repositories.CreateUserParams{
		Name:        params.Name,
		Email:       params.Email,
		Password:    encryptedPassword,
		PhoneNumber: params.PhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}