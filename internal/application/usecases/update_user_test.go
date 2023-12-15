package app_usecases

import (
	"errors"
	"testing"

	repositories_mock "github.com/sousair/americastech-user/internal/application/providers/repositories/mocks"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUserUseCase_UpdateRepositoryFindOneBy1Error(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	updateUserUC := NewUpdateUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, errors.New("error"))

	_, err := updateUserUC.Update(usecases.UpdateUserParams{})

	assert.Error(t, err)
}

func TestUpdateUserUseCase_UpdateUserNotFoundError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	updateUserUC := NewUpdateUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, nil)

	_, err := updateUserUC.Update(usecases.UpdateUserParams{})

	assert.Error(t, err)
}

func TestUpdateUserUseCase_UpdateEmailAlreadyExistsError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	updateUserUC := NewUpdateUserUseCase(userRepo)

	userRepo.On("FindOneBy", map[string]interface{}{
		"id": "1",
	}).Return(&entities.User{}, nil)

	userRepo.On("FindOneBy", map[string]interface{}{
		"email": "email",
	}).Return(&entities.User{}, nil)

	_, err := updateUserUC.Update(usecases.UpdateUserParams{
		ID:    "1",
		Email: "email",
	})

	assert.Error(t, err)
}

func TestUpdateUserUseCase_UpdateRepositoryFindOneBy2Error(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	updateUserUC := NewUpdateUserUseCase(userRepo)

	userRepo.On("FindOneBy", map[string]interface{}{
		"id": "1",
	}).Return(&entities.User{}, nil)

	userRepo.On("FindOneBy", map[string]interface{}{
		"email": "email",
	}).Return(nil, errors.New("error"))

	_, err := updateUserUC.Update(usecases.UpdateUserParams{
		ID:    "1",
		Email: "email",
	})

	assert.Error(t, err)
}
func TestUpdateUserUseCase_UpdateRepositoryUpdateError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	updateUserUC := NewUpdateUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	userRepo.On("Update", mock.Anything).Return(nil, errors.New("error"))

	_, err := updateUserUC.Update(usecases.UpdateUserParams{})

	assert.Error(t, err)
}

func TestUpdateUserUseCase_UpdateSuccess(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)

	updateUserUC := NewUpdateUserUseCase(userRepo)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	userRepo.On("Update", mock.Anything).Return(&entities.User{}, nil)

	user, err := updateUserUC.Update(usecases.UpdateUserParams{})

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
