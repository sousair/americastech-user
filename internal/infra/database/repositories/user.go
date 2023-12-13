package gorm_repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/entities"
	gorm_models "github.com/sousair/americastech-user/internal/infra/database/models"
	"gorm.io/gorm"
)

type (
	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(params repositories.CreateUserParams) (*entities.User, error) {
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

func (r UserRepository) FindAll() ([]*entities.User, error) {
	users := make([]*gorm_models.User, 0)

	if err := r.db.Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var usersEntities []*entities.User

	for _, user := range users {
		usersEntities = append(usersEntities, user.ToEntity())
	}

	return usersEntities, nil
}

func (r UserRepository) FindOneBy(where map[string]interface{}) (*entities.User, error) {
	userModel := &gorm_models.User{}

	if err := r.db.Where(where).First(userModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return userModel.ToEntity(), nil
}

func (r UserRepository) Update(user *entities.User) (*entities.User, error) {
	userModel := &gorm_models.User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		UpdatedAt:   time.Now(),
	}

	if err := r.db.Model(userModel).Updates(userModel).Error; err != nil {
		return nil, err
	}

	return r.FindOneBy(map[string]interface{}{
		"id": userModel.ID,
	})
}

func (r UserRepository) Delete(id string) error {
	if err := r.db.Unscoped().Delete(&gorm_models.User{ID: id}).Error; err != nil {
		return err
	}

	return nil
}
