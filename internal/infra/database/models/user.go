package gorm_models

import (
	"time"

	"github.com/sousair/americastech-user/internal/core/entities"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		ID          string    `gorm:"type:uuid;primaryKey"`
		Name        string    `gorm:"unique;not null"`
		Email       string    `gorm:"not null;uniqueIndex"`
		Password    string    `gorm:"not null"`
		PhoneNumber string    `gorm:"not null;unique"`
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
