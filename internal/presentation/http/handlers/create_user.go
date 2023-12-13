package http_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	CreateUserRequest struct {
		Name                 string `json:"name"`
		Email                string `json:"email"`
		Password             string `json:"password"`
		ConfirmationPassword string `json:"confirmation_password"`
		PhoneNumber          string `json:"phone_number"`
	}

	CreateUserResponse struct {
		User *entities.SanitizedUser `json:"user"`
	}

	createUserHandler struct {
		createUserUC usecases.CreateUserUseCase
	}
)

func NewCreateUserHandler(createUserUC usecases.CreateUserUseCase) *createUserHandler {
	return &createUserHandler{
		createUserUC: createUserUC,
	}
}

func (h *createUserHandler) Handle(c echo.Context) error {
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

	user, err := h.createUserUC.Create(usecases.CreateUserParams{
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

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": custom_errors.InternalServerError.Error(),
		})
	}

	return c.JSON(http.StatusCreated, CreateUserResponse{
		User: user.Sanitize(),
	})
}
