package app_usecases

import (
	"fmt"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/application/providers/cipher"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type createUserUseCase struct {
	userRepository repositories.UserRepository
	cipherProvider cipher.CipherProvider
}

func NewCreateUserUseCase(userRepo repositories.UserRepository, cipherProvider cipher.CipherProvider) usecases.CreateUserUseCase {
	return &createUserUseCase{
		userRepository: userRepo,
		cipherProvider: cipherProvider,
	}
}

func (uc createUserUseCase) Create(params usecases.CreateUserParams) (*entities.User, error) {
	emailAlreadyExists, err := uc.userRepository.FindOneBy(map[string]interface{}{
		"email": params.Email,
	})

	if err != nil {
		return nil, err
	}

	if emailAlreadyExists != nil {
		return nil, custom_errors.NewEmailAlreadyExistsError(
			fmt.Errorf("email %s already registered by user_id: %s", params.Email, emailAlreadyExists.ID),
			params.Email,
		)
	}

	encryptedPassword, err := uc.cipherProvider.Hash(params.Password)

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
