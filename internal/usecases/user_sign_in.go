package usecases

import (
	"time"

	custom_errors "github.com/sousair/americastech-user/internal/errors"
	"github.com/sousair/americastech-user/internal/providers/cryptography"
	"github.com/sousair/americastech-user/internal/providers/repositories"
)

type (
	UserSignInParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	TokenPayload struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserSignInResponse struct {
		Token   string       `json:"token"`
		Payload TokenPayload `json:"payload"`
	}

	UserSignInUseCase interface {
		SignIn(params UserSignInParams) (response *UserSignInResponse, err error)
	}

	userSignInUseCase struct {
		UserRepository repositories.UserRepository
		CryptoProvider cryptography.CryptoProvider
	}
)

func NewUserSignInUseCase(userRepo repositories.UserRepository, cryptoProvider cryptography.CryptoProvider) UserSignInUseCase {
	return &userSignInUseCase{
		UserRepository: userRepo,
		CryptoProvider: cryptoProvider,
	}
}

func (uc userSignInUseCase) SignIn(params UserSignInParams) (response *UserSignInResponse, err error) {
	user, err := uc.UserRepository.FindByEmail(params.Email)

	if err != nil {
		return nil, custom_errors.NewUserNotFoundError(err)
	}

	if err := uc.CryptoProvider.Compare(user.Password, params.Password); err != nil {
		return nil, custom_errors.NewInvalidPasswordError(err)
	}

	token, err := uc.CryptoProvider.GenerateAuthToken(cryptography.GenerateAuthTokenParams{
		Payload: map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		ExpirationTime: time.Now().Add(time.Hour * 3),
	})

	if err != nil {
		return nil, custom_errors.NewInternalServerError(err)
	}

	response = &UserSignInResponse{
		Token: token,
		Payload: TokenPayload{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil
}
