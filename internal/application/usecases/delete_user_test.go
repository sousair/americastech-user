package app_usecases

import (
	"errors"
	"testing"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	repositories_mock "github.com/sousair/americastech-user/internal/application/providers/repositories/mocks"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteUserUseCase_DeleteRepositoryFindOneByError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	deleteUserUC := NewDeleteUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, errors.New("error"))

	err := deleteUserUC.Delete("")

	assert.Error(t, err)
}

func TestDeleteUserUseCase_DeleteUserNotFoundError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	deleteUserUC := NewDeleteUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, nil)

	err := deleteUserUC.Delete("")

	assert.Error(t, err)
	assert.IsType(t, custom_errors.UserNotFoundError, err)
}

func TestDeleteUserUseCase_DeleteRepositoryDeleteError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	deleteUserUC := NewDeleteUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	userRepo.On("Delete", mock.Anything).Return(errors.New("error"))

	err := deleteUserUC.Delete("")

	assert.Error(t, err)
}

func TestDeleteUserUseCase_DeleteSuccess(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	deleteUserUC := NewDeleteUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	userRepo.On("Delete", mock.Anything).Return(nil)

	err := deleteUserUC.Delete("")

	assert.NoError(t, err)
}
