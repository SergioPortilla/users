package usescases

import (
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/domain/ports"
)

type GetByDniUseCase interface {
	Handler(DNI int16) ([]model.User, error)
}
type UseCaseGetByDni struct {
	UserRepository ports.UsersRepository
}

func (useCaseGetByDni *UseCaseGetByDni) Handler(DNI int16) ([]model.User, error) {
	return useCaseGetByDni.UserRepository.Get(DNI)
}
