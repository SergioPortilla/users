package usescases

import (
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/domain/ports"
)

type GetByDniUseCase interface {
	Handler(DNI int64) (*model.User, error)
}
type UseCaseGetByDni struct {
	UserRepository ports.UsersRepository
}

func (useCaseGetByDni *UseCaseGetByDni) Handler(DNI int64) (*model.User, error) {
	user, err := useCaseGetByDni.UserRepository.GetByDNI(DNI)
	if err != nil {
		return user, err
	}
	return user, nil
}
