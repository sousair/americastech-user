package app_usecases

import (
	"errors"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	updateUserUseCase struct {
		userRepository repositories.UserRepository
	}
)

func NewUpdateUserUseCase(userRepo repositories.UserRepository) usecases.UpdateUserUseCase {
	return &updateUserUseCase{
		userRepository: userRepo,
	}
}

func (uc updateUserUseCase) Update(params usecases.UpdateUserParams) (*entities.User, error) {
	user, err := uc.userRepository.FindOneBy(map[string]interface{}{
		"id": params.ID,
	})

	if err != nil {
		return nil, err

	}

	if user == nil {
		return nil, custom_errors.NewUserNotFoundError(
			errors.New("no user record found"),
		)
	}

	user.Name = params.Name
	user.Email = params.Email
	user.PhoneNumber = params.PhoneNumber

	user, err = uc.userRepository.Update(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
