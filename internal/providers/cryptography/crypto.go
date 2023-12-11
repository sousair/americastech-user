package cryptography

type (
	CryptoProvider interface {
		Hash(password string) (string, error)
	}
)
