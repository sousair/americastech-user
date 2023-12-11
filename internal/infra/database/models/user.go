package gorm_models

import (
	"time"

	"github.com/sousair/americastech-user/internal/entities"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		ID          string `gorm:"primaryKey"`
		Name        string
		Email       string `gorm:"uniqueIndex"`
		Password    string
		PhoneNumber string
		CreatedAt   time.Time `gorm:"autoCreateTime"`
		UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	}
)

func (u User) ToEntity() *entities.User {
	user := entities.User{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Password:    u.Password,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}

	return &user
}
