package usecases

import (
	"github.com/sousair/americastech-user/internal/entities"
	"github.com/sousair/americastech-user/internal/errors"
	"github.com/sousair/americastech-user/internal/repositories"
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
	}
)

var variableName int

func NewCreateUserUseCase(userRepo repositories.UserRepository) CreateUserUseCase {
	return &createUserUseCase{
		UserRepository: userRepo,
	}
}

func (uc *createUserUseCase) Create(params CreateUserParams) (*entities.User, error) {
	emailAlreadyExists, err := uc.UserRepository.FindByEmail(params.Email)

	if err != nil {
		return nil, err
	}

	if emailAlreadyExists != nil {
		return nil, &errors.EmailAlreadyExists{
			Email: params.Email,
		}
	}

	user, err := uc.UserRepository.Create(repositories.CreateUserParams{
		Name:        params.Name,
		Email:       params.Email,
		Password:    params.Password,
		PhoneNumber: params.PhoneNumber,
	})

	if err != nil {
		// TODO: Make a custom error for internal server error
		return nil, err
	}

	return user, nil
}
