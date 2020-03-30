package factory

import (
	"github.com/ceiba-meli-demo/users/application/commands"
	"github.com/ceiba-meli-demo/users/domain/model"
)

func CreateUser(userCommand commands.UserCommand) (model.User, error) {
	var user model.User
	user, err := user.CreateUser(userCommand.DNI, userCommand.Name, userCommand.LastName, userCommand.QuantityMovies)
	return user, err
}
