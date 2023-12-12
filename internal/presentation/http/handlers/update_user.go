package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-user/internal/entities"
	custom_errors "github.com/sousair/americastech-user/internal/errors"
	gorm_repositories "github.com/sousair/americastech-user/internal/infra/database/repositories"
	"gorm.io/gorm"
)

type (
	UpdateUserRequest struct {
		ID          string `param:"id"`
		Name        string `json:"name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}

	UpdateUserResponse struct {
		User *entities.SanitizedUser `json:"user"`
	}
)

func CreateUpdateUserHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var getUserRequest UpdateUserRequest

		if err := c.Bind(&getUserRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "invalid request body",
			})
		}

		authUserId := c.Get("user_id").(string)

		if authUserId != getUserRequest.ID {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "you are not allowed to access this resource",
			})
		}

		userRepo := gorm_repositories.NewUserRepository(db)

		user, err := userRepo.FindOneBy(map[string]interface{}{
			"id": getUserRequest.ID,
		})

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, map[string]string{
					"message": custom_errors.UserNotFoundError.Error(),
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": custom_errors.InternalServerError.Error(),
			})
		}

		user.Name = getUserRequest.Name
		user.Email = getUserRequest.Email
		user.PhoneNumber = getUserRequest.PhoneNumber

		user, err = userRepo.Update(user)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": custom_errors.InternalServerError.Error(),
			})
		}

		return c.JSON(http.StatusOK, UpdateUserResponse{
			User: user.Sanitize(),
		})
	}
}
