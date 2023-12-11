package http_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-user/internal/errors"
	gorm_repositories "github.com/sousair/americastech-user/internal/infra/database/repositories"
	"github.com/sousair/americastech-user/internal/usecases"
	"gorm.io/gorm"
)

type (
	CreateUserRequest struct {
		Name                 string `json:"name"`
		Email                string `json:"email"`
		Password             string `json:"password"`
		ConfirmationPassword string `json:"confirmation_password"`
		PhoneNumber          string `json:"phone_number"`
	}
)

func CreateUserHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var createUserRequest CreateUserRequest

		if err := c.Bind(&createUserRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				// TODO: Improve this message to be less generic
				"message": "invalid request body",
			})
		}

		if createUserRequest.Password != createUserRequest.ConfirmationPassword {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "password and confirmation_password does not match",
			})
		}

		userRepo := gorm_repositories.NewUserRepository(db)
		createUserUC := usecases.NewCreateUserUseCase(userRepo)

		user, err := createUserUC.Create(usecases.CreateUserParams{
			Name:        createUserRequest.Name,
			Email:       createUserRequest.Email,
			Password:    createUserRequest.Password,
			PhoneNumber: createUserRequest.PhoneNumber,
		})

		if err != nil {
			if errors.As(err, &custom_errors.EmailAlreadyExistsError) {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
			}

			if errors.As(err, &custom_errors.InternalServerError) {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"message": err.Error(),
				})
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"user": user,
		})
	}
}
