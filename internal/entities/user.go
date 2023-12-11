package entities

import "time"

type (
	User struct {
		ID          string    `json:"id"`
		Name        string    `json:"username"`
		Email       string    `json:"email"`
		Password    string    `json:"password"`
		PhoneNumber string    `json:"phone_number"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		DeletedAt   time.Time `json:"deleted_at"`
	}

	SanitizedUser struct {
		ID          string `json:"id"`
		Name        string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}
)

func (u User) Sanitize() *SanitizedUser {
	return &SanitizedUser{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}
