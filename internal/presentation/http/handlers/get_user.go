package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	gorm_repositories "github.com/sousair/americastech-user/internal/infra/database/repositories"
	"gorm.io/gorm"
)

type (
	GetUserRequest struct {
		ID string `param:"id"`
	}
)

func CreateGetUserHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var getUserRequest GetUserRequest

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
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		if user == nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "user not found",
			})
		}

		return c.JSON(http.StatusOK, user.Sanitize())
	}
}
