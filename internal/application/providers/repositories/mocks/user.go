package repositories_mock

import (
	"github.com/sousair/americastech-user/internal/application/providers/repositories"
	"github.com/sousair/americastech-user/internal/core/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(params repositories.CreateUserParams) (*entities.User, error) {
	args := m.Called(params)

	userR := args.Get(0)

	if userR == nil {
		return nil, args.Error(1)
	}

	return userR.(*entities.User), args.Error(1)
}

func (m *UserRepositoryMock) FindAll() ([]*entities.User, error) {
	args := m.Called()

	usersR := args.Get(0)

	if usersR == nil {
		return nil, args.Error(1)
	}

	return usersR.([]*entities.User), args.Error(1)
}

func (m *UserRepositoryMock) FindOneBy(params map[string]interface{}) (*entities.User, error) {
	args := m.Called(params)

	userR := args.Get(0)

	if userR == nil {
		return nil, args.Error(1)
	}

	return userR.(*entities.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(user *entities.User) (*entities.User, error) {
	args := m.Called(user)

	userR := args.Get(0)

	if userR == nil {
		return nil, args.Error(1)
	}

	return userR.(*entities.User), args.Error(1)
}

func (m *UserRepositoryMock) Delete(id string) error {
	args := m.Called(id)

	return args.Error(0)
}
