package repositories

import "github.com/sousair/americastech-user/internal/entities"

type (
	CreateUserParams struct {
		Name        string `json:"username"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		PhoneNumber string `json:"phone_number"`
	}

	UserRepository interface {
		Create(params CreateUserParams) (*entities.User, error)
		FindByEmail(email string) (*entities.User, error)
		FindAll() ([]*entities.User, error)
		FindOneBy(where map[string]interface{}) (*entities.User, error)
	}
)
