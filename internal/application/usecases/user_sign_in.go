package app_usecases

import (
	"time"

	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/application/providers/cryptography"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type userSignInUseCase struct {
	userRepository repositories.UserRepository
	cryptoProvider cryptography.CryptoProvider
}

func NewUserSignInUseCase(userRepo repositories.UserRepository, cryptoProvider cryptography.CryptoProvider) usecases.UserSignInUseCase {
	return &userSignInUseCase{
		userRepository: userRepo,
		cryptoProvider: cryptoProvider,
	}
}

func (uc userSignInUseCase) SignIn(params usecases.UserSignInParams) (response *usecases.UserSignInResponse, err error) {
	user, err := uc.userRepository.FindOneBy(map[string]interface{}{
		"email": params.Email,
	})

	if err != nil {
		return nil, custom_errors.NewUserNotFoundError(err)
	}

	if err := uc.cryptoProvider.Compare(user.Password, params.Password); err != nil {
		return nil, custom_errors.NewInvalidPasswordError(err)
	}

	token, err := uc.cryptoProvider.GenerateAuthToken(cryptography.GenerateAuthTokenParams{
		Payload: map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		ExpirationTime: time.Now().Add(time.Hour * 3),
	})

	if err != nil {
		return nil, err
	}

	response = &usecases.UserSignInResponse{
		Token: token,
		Payload: usecases.TokenPayload{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil
}