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
		UserRepository repositories.UserRepository
		CryptoProvider cryptography.CryptoProvider
	}
)

var variableName int

func NewCreateUserUseCase(userRepo repositories.UserRepository, cryptoProvider cryptography.CryptoProvider) CreateUserUseCase {
	return &createUserUseCase{
		UserRepository: userRepo,
		CryptoProvider: cryptoProvider,
	}
}

func (uc *createUserUseCase) Create(params CreateUserParams) (*entities.User, error) {
	emailAlreadyExists, _ := uc.UserRepository.FindByEmail(params.Email)

	if emailAlreadyExists != nil {
		return nil, custom_errors.NewEmailAlreadyExistsError(errors.New(""), params.Email)
	}

	encryptedPassword, err := uc.CryptoProvider.Hash(params.Password)

	if err != nil {
		return nil, custom_errors.NewInternalServerError(err)
	}

	user, err := uc.UserRepository.Create(repositories.CreateUserParams{
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
