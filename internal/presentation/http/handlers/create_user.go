package http_handlers

import (
	"github.com/labstack/echo/v4"
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
			return c.JSON(400, map[string]string{
				// TODO: Improve this message to be less generic
				"message": "Invalid request body",
			})
		}

		if createUserRequest.Password != createUserRequest.ConfirmationPassword {
			return c.JSON(400, map[string]string{
				"message": "Password and confirmation password does not match",
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
			// TODO: Improve this
			return c.JSON(400, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(200, map[string]interface{}{
			"user": user,
		})
	}
}
