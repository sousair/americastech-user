package http_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sousair/americastech-user/internal/core/entities"
	gorm_repositories "github.com/sousair/americastech-user/internal/infra/database/repositories"
	"gorm.io/gorm"
)

type (
	GetUsersResponse struct {
		Users []*entities.SanitizedUser `json:"users"`
	}
)

func CreateGetUsersHandler(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		userRepo := gorm_repositories.NewUserRepository(db)

		users, err := userRepo.FindAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		var sanitizedUsers []*entities.SanitizedUser

		for _, user := range users {
			sanitizedUsers = append(sanitizedUsers, user.Sanitize())
		}

		return c.JSON(http.StatusOK, GetUsersResponse{
			Users: sanitizedUsers,
		})
	}
}
