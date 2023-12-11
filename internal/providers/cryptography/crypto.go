package cryptography

import "time"

type (
	GenerateAuthTokenParams struct {
		Payload        map[string]interface{}
		ExpirationTime time.Time
	}

	CryptoProvider interface {
		Hash(password string) (string, error)
		Compare(hash string, password string) error
		GenerateAuthToken(GenerateAuthTokenParams) (string, error)
	}
)
