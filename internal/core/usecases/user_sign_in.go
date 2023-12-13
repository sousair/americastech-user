package usecases

type (
	UserSignInParams struct {
		Email    string
		Password string
	}

	UserSignInResponse struct {
		Token   string
		Payload TokenPayload
	}

	TokenPayload struct {
		ID    string
		Name  string
		Email string
	}

	UserSignInUseCase interface {
		SignIn(params UserSignInParams) (response *UserSignInResponse, err error)
	}
)
