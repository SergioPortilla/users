package usescases

import (
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/domain/ports"
)

type UpdateUserUseCase interface {
	Handler(DNI int64) (*model.Message, error)
}

type UseCaseUpdateUser struct {
	UserRepository ports.UsersRepository
}

func (useCaseUpdateUser *UseCaseUpdateUser) Handler(DNI int64) (*model.Message, error) {
	userUpdated, err := useCaseUpdateUser.UserRepository.UpdateQuantityMovies(DNI)
	if err != nil {
		return userUpdated, err
	}
	return userUpdated, nil
}
