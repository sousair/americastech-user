package http_handlers

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	UpdateUserRequest struct {
		ID          string `param:"id" validate:"required,uuid4"`
		Name        string `json:"name" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number" validate:"required"`
	}

	UpdateUserResponse struct {
		User *entities.SanitizedUser `json:"user"`
	}

	updateUserHandler struct {
		updateUserUC usecases.UpdateUserUseCase
		validator    *validator.Validate
	}
)

func NewUpdateUserHandler(updateUserUC usecases.UpdateUserUseCase, validator *validator.Validate) *updateUserHandler {
	return &updateUserHandler{
		updateUserUC: updateUserUC,
		validator:    validator,
	}
}

func (h *updateUserHandler) Handle(c echo.Context) error {
	var getUserRequest UpdateUserRequest

	if err := c.Bind(&getUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	if err := h.validator.Struct(getUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	authUserId := c.Get("user_id").(string)

	if authUserId != getUserRequest.ID {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "you are not allowed to access this resource",
		})
	}

	user, err := h.updateUserUC.Update(usecases.UpdateUserParams{
		ID:          getUserRequest.ID,
		Name:        getUserRequest.Name,
		Email:       getUserRequest.Email,
		PhoneNumber: getUserRequest.PhoneNumber,
	})

	if err != nil {
		if errors.As(err, &custom_errors.UserNotFoundError) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		}

		if errors.As(err, &custom_errors.EmailAlreadyExistsError) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": custom_errors.InternalServerError.Error(),
		})
	}

	return c.JSON(http.StatusOK, UpdateUserResponse{
		User: user.Sanitize(),
	})
}
