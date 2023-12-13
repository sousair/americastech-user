package jwt

import "time"

type (
	GenerateAuthTokenParams struct {
		Payload        map[string]interface{}
		ExpirationTime time.Time
	}

	JWTProvider interface {
		GenerateAuthToken(params GenerateAuthTokenParams) (string, error)
		ValidateAuthToken(token string) (tokenPayload map[string]interface{}, err error)
	}
)
