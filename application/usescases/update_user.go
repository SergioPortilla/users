package usescases

import (
	"github.com/ceiba-meli-demo/users/application/commands"
	"github.com/ceiba-meli-demo/users/application/factory"
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/domain/ports"
)

type UpdateUserUseCase interface {
	Handler(DNI int64, userCommand commands.UserCommand) (*model.User, error)
}

type UseCaseUpdateUser struct {
	UserRepository ports.UsersRepository
}

func (useCaseUpdateUser *UseCaseUpdateUser) Handler(DNI int64, userCommand commands.UserCommand) (*model.User, error) {
	user, err := factory.CreateUser(userCommand)
	if err != nil {
		return nil, err
	}
	userUpdated, err := useCaseUpdateUser.UserRepository.UpdateQuantityMovies(DNI, user)
	if err != nil {
		return userUpdated, err
	}
	return userUpdated, nil
}
