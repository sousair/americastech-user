package repositories

import "github.com/sousair/americastech-user/internal/entities"

type (
	CreateUserParams struct {
		Name        string
		Email       string
		Password    string
		PhoneNumber string
	}

	UserRepository interface {
		Create(params CreateUserParams) (*entities.User, error)
		FindAll() ([]*entities.User, error)
		FindOneBy(where map[string]interface{}) (*entities.User, error)
		Update(user *entities.User) (*entities.User, error)
	}
)
