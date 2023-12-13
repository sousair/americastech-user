package http_handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-user/internal/application/errors"
	"github.com/sousair/americastech-user/internal/core/usecases"
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

	userSignInHandler struct {
		userSignInUC usecases.UserSignInUseCase
	}
)

func NewUserSignInHandler(userSignInUC usecases.UserSignInUseCase) *userSignInHandler {
	return &userSignInHandler{
		userSignInUC: userSignInUC,
	}
}

func (h *userSignInHandler) Handle(c echo.Context) error {
	var userSignInRequest UserSignInRequest

	if err := c.Bind(&userSignInRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			// TODO: Improve this message to be less generic
			"message": "invalid request body",
		})
	}

	signResponse, err := h.userSignInUC.SignIn(usecases.UserSignInParams{
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

	return c.JSON(http.StatusOK, &UserSignInResponse{
		Token:   signResponse.Token,
		Payload: signResponse.Payload,
	})
}
