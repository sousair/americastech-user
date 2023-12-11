package crypto_provider

import "golang.org/x/crypto/bcrypt"

type (
	CryptoProvider struct{}
)

func NewCryptoProvider() *CryptoProvider {
	return &CryptoProvider{}
}

func (cp *CryptoProvider) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
