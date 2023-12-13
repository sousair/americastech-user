package http_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	DeleteUserRequest struct {
		ID string `param:"id"`
	}

	deleteUserHandler struct {
		deleteUserUC usecases.DeleteUserUseCase
	}
)

func NewDeleteUserHandler(deleteUserUC usecases.DeleteUserUseCase) *deleteUserHandler {
	return &deleteUserHandler{
		deleteUserUC: deleteUserUC,
	}
}

func (h *deleteUserHandler) Handle(c echo.Context) error {
	var deleteUserRequest DeleteUserRequest

	if err := c.Bind(&deleteUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request body",
		})
	}

	authUserId := c.Get("user_id").(string)

	if authUserId != deleteUserRequest.ID {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "you are not allowed to access this resource",
		})
	}

	err := h.deleteUserUC.Delete(deleteUserRequest.ID)

	if err != nil {
		if errors.As(err, &custom_errors.UserNotFoundError) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": custom_errors.UserNotFoundError.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": custom_errors.InternalServerError.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user deleted successfully",
	})
}
