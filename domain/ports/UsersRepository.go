package ports

import "github.com/ceiba-meli-demo/users/domain/model"

type UsersRepository interface {
	Get(DNI int16) (model.User, error)
	Update(DNI int16, user model.User) (*model.User, error)
}
