package usecases

type (
	UserSignInParams struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserSignInResponse struct {
		Token   string       `json:"token"`
		Payload TokenPayload `json:"payload"`
	}

	TokenPayload struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserSignInUseCase interface {
		SignIn(params UserSignInParams) (response *UserSignInResponse, err error)
	}
)
