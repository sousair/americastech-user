package crypto_provider

import (
	"os"

	"github.com/sousair/americastech-user/internal/providers/cryptography"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

type (
	CryptoProvider struct {
		userTokenSecret string
	}
)

func NewCryptoProvider() cryptography.CryptoProvider {
	userTokenSecret := os.Getenv("USER_TOKEN_SECRET")
	return &CryptoProvider{
		userTokenSecret: userTokenSecret,
	}
}

func (cp CryptoProvider) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (cp CryptoProvider) Compare(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (cp CryptoProvider) GenerateAuthToken(params cryptography.GenerateAuthTokenParams) (string, error) {
	userPayload := jwt.MapClaims{
		"id":    params.Payload["id"],
		"name":  params.Payload["name"],
		"email": params.Payload["email"],
		"exp":   params.ExpirationTime.Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(userPayload)).SignedString([]byte(cp.userTokenSecret))
}
