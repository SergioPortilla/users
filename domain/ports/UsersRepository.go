package ports

import "github.com/ceiba-meli-demo/users/domain/model"

type UsersRepository interface {
	GetByDNI(DNI int64) (*model.User, error)
	UpdateQuantityMovies(DNI int64) (*model.Message, error)
}
