package app_usecases

import (
	"errors"
	"testing"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	cipher_mock "github.com/sousair/americastech-user/internal/application/providers/cipher/mocks"
	repositories_mock "github.com/sousair/americastech-user/internal/application/providers/repositories/mocks"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserUseCase_CreateRepositoryFindOneByError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)

	createUserUC := NewCreateUserUseCase(userRepo, cipherProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, errors.New("error"))

	_, err := createUserUC.Create(usecases.CreateUserParams{})

	assert.Error(t, err)
}
func TestCreateUserUseCase_CreateEmailAlreadyRegisteredError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)

	createUserUC := NewCreateUserUseCase(userRepo, cipherProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)

	_, err := createUserUC.Create(usecases.CreateUserParams{})

	assert.Error(t, err)
	assert.IsType(t, custom_errors.EmailAlreadyExistsError, err)
}

func TestCreateUserUseCase_CreateCipherHashError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)

	createUserUC := NewCreateUserUseCase(userRepo, cipherProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, nil)
	cipherProvider.On("Hash", mock.Anything).Return("", errors.New("error"))

	_, err := createUserUC.Create(usecases.CreateUserParams{})

	assert.Error(t, err)
}

func TestCreateUserUseCase_CreateRepositoryCreateError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)

	createUserUC := NewCreateUserUseCase(userRepo, cipherProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, nil)
	cipherProvider.On("Hash", mock.Anything).Return("hash", nil)
	userRepo.On("Create", mock.Anything).Return(nil, errors.New("error"))

	_, err := createUserUC.Create(usecases.CreateUserParams{})

	assert.Error(t, err)
}

func TestCreateUserUseCase_CreateSuccess(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)

	createUserUC := NewCreateUserUseCase(userRepo, cipherProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, nil)
	cipherProvider.On("Hash", mock.Anything).Return("hash", nil)
	userRepo.On("Create", mock.Anything).Return(&entities.User{}, nil)

	user, err := createUserUC.Create(usecases.CreateUserParams{})

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
