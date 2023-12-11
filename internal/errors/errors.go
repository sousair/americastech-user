package custom_errors

import "fmt"

// TODO: Break this in multiple files
type emailAlreadyExistsError struct {
	email string
	cause error
}

var EmailAlreadyExistsError = &emailAlreadyExistsError{}

func NewEmailAlreadyExistsError(err error, email string) *emailAlreadyExistsError {
	return &emailAlreadyExistsError{
		email: email,
		cause: err,
	}
}

func (e emailAlreadyExistsError) Error() string {
	return fmt.Sprintf("email %v already exists", e.email)
}

type internalServerError struct {
	cause error
}

var InternalServerError = &internalServerError{}

func NewInternalServerError(cause error) *internalServerError {
	return &internalServerError{
		cause: cause,
	}
}

func (e internalServerError) Error() string {
	return "internal server error"
}
