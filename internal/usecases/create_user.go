package usecases

import (
	"errors"

	"github.com/sousair/americastech-user/internal/entities"
	custom_errors "github.com/sousair/americastech-user/internal/errors"
	"github.com/sousair/americastech-user/internal/providers/cryptography"
	"github.com/sousair/americastech-user/internal/providers/repositories"
)

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

	createUserUseCase struct {
		userRepository repositories.UserRepository
		cryptoProvider cryptography.CryptoProvider
	}
)

func NewCreateUserUseCase(userRepo repositories.UserRepository, cryptoProvider cryptography.CryptoProvider) CreateUserUseCase {
	return &createUserUseCase{
		userRepository: userRepo,
		cryptoProvider: cryptoProvider,
	}
}

func (uc createUserUseCase) Create(params CreateUserParams) (*entities.User, error) {
	emailAlreadyExists, _ := uc.userRepository.FindOneBy(map[string]interface{}{
		"email": params.Email,
	})

	if emailAlreadyExists != nil {
		return nil, custom_errors.NewEmailAlreadyExistsError(errors.New(""), params.Email)
	}

	encryptedPassword, err := uc.cryptoProvider.Hash(params.Password)

	if err != nil {
		return nil, custom_errors.NewInternalServerError(err)
	}

	user, err := uc.userRepository.Create(repositories.CreateUserParams{
		Name:        params.Name,
		Email:       params.Email,
		Password:    encryptedPassword,
		PhoneNumber: params.PhoneNumber,
	})

	if err != nil {
		return nil, custom_errors.NewInternalServerError(err)
	}

	return user, nil
}
