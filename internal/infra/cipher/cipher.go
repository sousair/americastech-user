package bcrypt_cipher

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/sousair/americastech-user/internal/application/providers/cipher"
)

type bcryptCipherProvider struct {
	userPasswordCost int
}

func NewCipherProvider(userPasswordCost int) cipher.CipherProvider {
	return &bcryptCipherProvider{
		userPasswordCost: userPasswordCost,
	}
}

func (p bcryptCipherProvider) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), p.userPasswordCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (p bcryptCipherProvider) Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
