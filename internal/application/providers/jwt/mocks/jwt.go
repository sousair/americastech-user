package jwt_mock

import (
	"github.com/sousair/americastech-user/internal/application/providers/jwt"
	"github.com/stretchr/testify/mock"
)

type JWTProviderMock struct {
	mock.Mock
}

func (m *JWTProviderMock) GenerateAuthToken(params jwt.GenerateAuthTokenParams) (string, error) {
	args := m.Called(params)

	return args.String(0), args.Error(1)
}

func (m *JWTProviderMock) ValidateAuthToken(token string) (tokenPayload map[string]interface{}, err error) {
	args := m.Called(token)

	tokenPayloadR := args.Get(0)

	if tokenPayloadR == nil {
		return nil, args.Error(1)
	}

	return tokenPayloadR.(map[string]interface{}), args.Error(1)
}
