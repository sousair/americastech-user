package usecases

type (
	DeleteUserUseCase interface {
		Delete(id string) error
	}
)
