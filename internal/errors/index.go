package errors

import "fmt"

type (
	EmailAlreadyExists struct {
		Email string
	}
)

func (e *EmailAlreadyExists) Error() string {
	return fmt.Sprintf("email %v already exists", e.Email)
}
