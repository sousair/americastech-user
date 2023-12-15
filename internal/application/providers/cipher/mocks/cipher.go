package cipher_mock

import "github.com/stretchr/testify/mock"

type CipherProviderMock struct {
	mock.Mock
}

func (m *CipherProviderMock) Hash(password string) (string, error) {
	args := m.Called(password)

	return args.String(0), args.Error(1)

}

func (m *CipherProviderMock) Compare(hash string, password string) error {
	args := m.Called(hash, password)

	return args.Error(0)
}
