package gorm_repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/sousair/americastech-user/internal/entities"
	gorm_models "github.com/sousair/americastech-user/internal/infra/database/models"
	"github.com/sousair/americastech-user/internal/repositories"
	"gorm.io/gorm"
)

type (
	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(params repositories.CreateUserParams) (*entities.User, error) {
	user := &gorm_models.User{
		ID:          uuid.New().String(),
		Name:        params.Name,
		Email:       params.Email,
		Password:    params.Password,
		PhoneNumber: params.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}

func (r *UserRepository) FindByEmail(email string) (*entities.User, error) {
	user := &gorm_models.User{}

	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}
