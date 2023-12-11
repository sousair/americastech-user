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

type userNotFoundError struct {
	cause error
}

var UserNotFoundError = &userNotFoundError{}

func NewUserNotFoundError(cause error) *userNotFoundError {
	return &userNotFoundError{
		cause: cause,
	}
}

func (e userNotFoundError) Error() string {
	return "user not found"
}

type invalidPasswordError struct {
	cause error
}

var InvalidPasswordError = &invalidPasswordError{}

func NewInvalidPasswordError(cause error) *invalidPasswordError {
	return &invalidPasswordError{
		cause: cause,
	}
}

func (e invalidPasswordError) Error() string {
	return "invalid password"
}
