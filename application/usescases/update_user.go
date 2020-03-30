package usescases

import (
	"github.com/ceiba-meli-demo/users/application/commands"
	"github.com/ceiba-meli-demo/users/domain/factory"
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/domain/ports"
)

type UpdateUserUseCase interface {
	Handler(DNI int16, userCommand commands.UserCommand) (*model.User, error)
}

type UseCaseUpdateUser struct {
	UserRepository ports.UsersRepository
}

func (useCaseUpdateUser *UseCaseUpdateUser) Handler(DNI int16, userCommand commands.UserCommand) (*model.User, error) {
	user, err := factory.CreateUser(userCommand)
	if err != nil {
		return nil, err
	}
	userUpdated, err := useCaseUpdateUser.UserRepository.Update(DNI, user)
	if err != nil {
		return userUpdated, err
	}
	return userUpdated, nil
}
