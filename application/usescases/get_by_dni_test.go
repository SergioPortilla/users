package usescases

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type useCaseGetByDniMock struct {
	mock.Mock
}

func (m *useCaseGetByDniMock) GetByDni(DNI int64) bool {
	args := m.Called(DNI)
	return args.Bool(0)
}

func TestGetByDni(t *testing.T) {
	getUserByDni := new(useCaseGetByDniMock)

	getUserByDni.On("GetByDni", 1031142378).Return(true)

	getByDniUseCase := GetByDniUseCase{getUserByDni}

	getByDniUseCase.Handler(1031142378)

	getUserByDni.AssertExpectations(t)

}
