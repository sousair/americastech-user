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
	}

	SanitizedUser struct {
		ID          string `json:"id"`
		Name        string `json:"username"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
)

func (u User) Sanitize() *SanitizedUser {
	return &SanitizedUser{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   u.UpdatedAt.Format(time.RFC3339),
	}
}
