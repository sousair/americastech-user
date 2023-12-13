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
	GetUsersResponse struct {
		Users []*entities.SanitizedUser `json:"users"`
	}

	getUsersHandler struct {
		getUsersUC usecases.GetUsersUseCase
	}
)

func NewGetUsersHandler(getUsersUC usecases.GetUsersUseCase) *getUsersHandler {
	return &getUsersHandler{
		getUsersUC: getUsersUC,
	}
}

func (h *getUsersHandler) Handle(c echo.Context) error {
	users, err := h.getUsersUC.Get()

	if err != nil {
		if errors.As(err, &custom_errors.UserNotFoundError) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	var sanitizedUsers []*entities.SanitizedUser

	for _, user := range users {
		sanitizedUsers = append(sanitizedUsers, user.Sanitize())
	}

	return c.JSON(200, GetUsersResponse{
		Users: sanitizedUsers,
	})
}
