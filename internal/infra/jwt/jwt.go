package jwt_provider

import (
	jwt_go "github.com/dgrijalva/jwt-go"

	"github.com/sousair/americastech-user/internal/application/providers/jwt"
	"github.com/sousair/americastech-user/internal/core/usecases"
)

type (
	jwtProvider struct {
		tokenSecret string
	}

	UserTokenClaims struct {
		usecases.TokenPayload
		jwt_go.StandardClaims
	}
)

func NewJwtProvider(tokenSecret string) jwt.JWTProvider {
	return &jwtProvider{
		tokenSecret: tokenSecret,
	}
}

func (p jwtProvider) GenerateAuthToken(params jwt.GenerateAuthTokenParams) (string, error) {
	userPayload := jwt_go.MapClaims{
		"id":    params.Payload["id"],
		"name":  params.Payload["name"],
		"email": params.Payload["email"],
		"exp":   params.ExpirationTime.Unix(),
	}

	return jwt_go.NewWithClaims(jwt_go.SigningMethodHS256, userPayload).SignedString([]byte(p.tokenSecret))
}

func (p jwtProvider) ValidateAuthToken(token string) (payload map[string]interface{}, err error) {
	claims := &UserTokenClaims{}
	parsedToken, err := jwt_go.ParseWithClaims(token, claims, func(token *jwt_go.Token) (interface{}, error) {
		return []byte(p.tokenSecret), nil
	})

	if err != nil {
		if err == jwt_go.ErrSignatureInvalid {
			return nil, err
		}

		return nil, err
	}

	if !parsedToken.Valid {
		return nil, jwt_go.ErrInvalidKey
	}

	return map[string]interface{}{
		"id":    claims.ID,
		"name":  claims.Name,
		"email": claims.Email,
	}, nil
}
