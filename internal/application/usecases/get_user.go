package app_usecases

import (
	"errors"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type getUserUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetUserUseCase(userRepo repositories.UserRepository) usecases.GetUserUseCase {
	return &getUserUseCase{
		userRepository: userRepo,
	}
}

func (uc getUserUseCase) Get(params usecases.GetUserParams) (*entities.User, error) {
	user, err := uc.userRepository.FindOneBy(map[string]interface{}{
		"id": params.ID,
	})

	if user == nil {
		return nil, err
	}

	if err != nil {
		return nil, custom_errors.NewUserNotFoundError(
			errors.New("no user record found"),
		)
	}

	return user, nil
}
