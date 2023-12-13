package app_usecases

import (
	"errors"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
)

type (
	deleteUserUseCase struct {
		userRepository repositories.UserRepository
	}
)

func NewDeleteUserUseCase(userRepo repositories.UserRepository) *deleteUserUseCase {
	return &deleteUserUseCase{
		userRepository: userRepo,
	}
}

func (uc deleteUserUseCase) Delete(id string) error {
	user, err := uc.userRepository.FindOneBy(map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	if user == nil {
		return custom_errors.NewUserNotFoundError(
			errors.New("no user record found"),
		)
	}

	if err := uc.userRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
