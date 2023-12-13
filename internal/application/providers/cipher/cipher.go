package cipher

type (
	CipherProvider interface {
		Hash(password string) (string, error)
		Compare(hash string, password string) error
	}
)
