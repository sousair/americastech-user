package crypto_provider

import (
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/sousair/americastech-user/internal/application/providers/cryptography"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	CryptoProvider struct {
		userTokenSecret string
	}

	UserTokenClaims struct {
		usecases.TokenPayload
		jwt.StandardClaims
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

	return jwt.NewWithClaims(jwt.SigningMethodHS256, userPayload).SignedString([]byte(cp.userTokenSecret))
}

func (cp CryptoProvider) VerifyAuthToken(token string) (payload map[string]interface{}, err error) {
	claims := &UserTokenClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cp.userTokenSecret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}

		return nil, err
	}

	if !parsedToken.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return map[string]interface{}{
		"id":    claims.ID,
		"name":  claims.Name,
		"email": claims.Email,
	}, nil
}
