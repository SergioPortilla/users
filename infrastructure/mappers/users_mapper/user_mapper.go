package users_mapper

import (
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/models"
)

func UserToUserDb(user model.User) models.UserDb {

	var userDb models.UserDb
	userDb.DNI = user.DNI
	userDb.Name = user.Name
	userDb.LastName = user.LastName
	userDb.QuantityMovies = user.QuantityMovies

	return userDb

}

func UserDbToUser(userDb models.UserDb) model.User {

	var user model.User
	user.DNI = userDb.DNI
	user.Name = userDb.Name
	user.LastName = userDb.LastName
	user.QuantityMovies = userDb.QuantityMovies

	return user

}
