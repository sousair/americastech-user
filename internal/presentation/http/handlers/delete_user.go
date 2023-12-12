package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	custom_errors "github.com/sousair/americastech-user/internal/errors"
	gorm_repositories "github.com/sousair/americastech-user/internal/infra/database/repositories"
	"gorm.io/gorm"
)

type (
	DeleteUserRequest struct {
		ID string `param:"id"`
	}
)

func CreateDeleteUserHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
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

		userRepo := gorm_repositories.NewUserRepository(db)

		err := userRepo.Delete(deleteUserRequest.ID)

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

		return c.JSON(http.StatusOK, map[string]string{
			"message": "user deleted successfully",
		})
	}
}
