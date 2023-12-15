package app_usecases

import (
	"errors"
	"testing"

	cipher_mock "github.com/sousair/americastech-user/internal/application/providers/cipher/mocks"
	jwt_mock "github.com/sousair/americastech-user/internal/application/providers/jwt/mocks"
	repositories_mock "github.com/sousair/americastech-user/internal/application/providers/repositories/mocks"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserSignInUseCase_SignInRepositoryFindOneError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)
	jwtProvider := new(jwt_mock.JWTProviderMock)

	userSignInUC := NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, errors.New("error"))

	_, err := userSignInUC.SignIn(usecases.UserSignInParams{})

	assert.Error(t, err)
}

func TestUserSignInUseCase_SignInUserNotFoundError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)
	jwtProvider := new(jwt_mock.JWTProviderMock)

	userSignInUC := NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(nil, nil)

	_, err := userSignInUC.SignIn(usecases.UserSignInParams{})

	assert.Error(t, err)
}

func TestUserSignInUseCase_SignInInvalidPasswordError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)
	jwtProvider := new(jwt_mock.JWTProviderMock)

	userSignInUC := NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	cipherProvider.On("Compare", mock.Anything, mock.Anything).Return(errors.New("error"))

	_, err := userSignInUC.SignIn(usecases.UserSignInParams{})

	assert.Error(t, err)
}

func TestUserSignInUseCase_SignInJWTProviderGenerateAuthTokenError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)
	jwtProvider := new(jwt_mock.JWTProviderMock)

	userSignInUC := NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	cipherProvider.On("Compare", mock.Anything, mock.Anything).Return(nil)
	jwtProvider.On("GenerateAuthToken", mock.Anything).Return("", errors.New("error"))

	_, err := userSignInUC.SignIn(usecases.UserSignInParams{})

	assert.Error(t, err)
}

func TestUserSignInUseCase_SignInEmptyTokenGeneratedError(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)
	jwtProvider := new(jwt_mock.JWTProviderMock)

	userSignInUC := NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	cipherProvider.On("Compare", mock.Anything, mock.Anything).Return(nil)
	jwtProvider.On("GenerateAuthToken", mock.Anything).Return("", nil)

	_, err := userSignInUC.SignIn(usecases.UserSignInParams{})

	assert.Error(t, err)
}

func TestUserSignInUseCase_SignInSuccess(t *testing.T) {
	userRepo := new(repositories_mock.UserRepositoryMock)
	cipherProvider := new(cipher_mock.CipherProviderMock)
	jwtProvider := new(jwt_mock.JWTProviderMock)

	userSignInUC := NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)

	userRepo.On("FindOneBy", mock.Anything).Return(&entities.User{}, nil)
	cipherProvider.On("Compare", mock.Anything, mock.Anything).Return(nil)
	jwtProvider.On("GenerateAuthToken", mock.Anything).Return("token", nil)

	signInResponse, err := userSignInUC.SignIn(usecases.UserSignInParams{})

	assert.NoError(t, err)
	assert.NotNil(t, signInResponse)
}
