package app_usecases

import (
	"errors"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	getUsersUseCase struct {
		userRepository repositories.UserRepository
	}
)

func NewGetUsersUseCase(userRepo repositories.UserRepository) usecases.GetUsersUseCase {
	return &getUsersUseCase{
		userRepository: userRepo,
	}
}

func (uc getUsersUseCase) Get() ([]*entities.User, error) {
	users, err := uc.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, custom_errors.NewUserNotFoundError(
			errors.New("no users record found"),
		)
	}

	return users, nil
}
