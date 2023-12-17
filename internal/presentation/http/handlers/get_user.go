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
	GetUserRequest struct {
		ID string `param:"id" validate:"required,uuid4"`
	}

	GetUserResponse struct {
		User *entities.SanitizedUser `json:"user"`
	}

	getUserHandler struct {
		getUserUC usecases.GetUserUseCase
		validator *validator.Validate
	}
)

func NewGetUserHandler(getUserUC usecases.GetUserUseCase, validator *validator.Validate) *getUserHandler {
	return &getUserHandler{
		getUserUC: getUserUC,
		validator: validator,
	}
}

func (h *getUserHandler) Handle(c echo.Context) error {
	var getUserRequest GetUserRequest

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

	user, err := h.getUserUC.Get(usecases.GetUserParams{
		ID: getUserRequest.ID,
	})

	if err != nil {
		if errors.As(err, &custom_errors.UserNotFoundError) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": custom_errors.EmailAlreadyExistsError.Error(),
		})
	}

	return c.JSON(http.StatusOK, GetUserResponse{
		User: user.Sanitize(),
	})
}
