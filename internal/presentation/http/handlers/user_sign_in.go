package http_handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	app_usecases "github.com/sousair/americastech-user/internal/application/usecases"
	"github.com/sousair/americastech-user/internal/core/usecases"
	bcrypt_cipher "github.com/sousair/americastech-user/internal/infra/cipher"
	gorm_repositories "github.com/sousair/americastech-user/internal/infra/database/repositories"
	jwt_provider "github.com/sousair/americastech-user/internal/infra/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	UserSignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserSignInResponse struct {
		Token   string                `json:"token"`
		Payload usecases.TokenPayload `json:"payload"`
	}
)

func CreateUserSignInHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var userSignInRequest UserSignInRequest

		if err := c.Bind(&userSignInRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				// TODO: Improve this message to be less generic
				"message": "invalid request body",
			})
		}

		userRepo := gorm_repositories.NewUserRepository(db)
		// Get cost from env.
		cipherProvider := bcrypt_cipher.NewCipherProvider(bcrypt.DefaultCost)
		userSecret := os.Getenv("USER_TOKEN_SECRET")
		jwtProvider := jwt_provider.NewJwtProvider(userSecret)

		userSignInUC := app_usecases.NewUserSignInUseCase(userRepo, cipherProvider, jwtProvider)

		signResponse, err := userSignInUC.SignIn(usecases.UserSignInParams{
			Email:    userSignInRequest.Email,
			Password: userSignInRequest.Password,
		})

		if err != nil {
			if errors.As(err, &custom_errors.UserNotFoundError) {
				return c.JSON(http.StatusNotFound, map[string]string{
					"message": err.Error(),
				})
			}

			if errors.As(err, &custom_errors.InvalidPasswordError) {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": err.Error(),
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, signResponse)
	}
}
