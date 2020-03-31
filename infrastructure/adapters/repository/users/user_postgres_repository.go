package users

import (
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/models"
	"github.com/ceiba-meli-demo/users/infrastructure/mappers/users_mapper"
	"github.com/jinzhu/gorm"
)

type UserPostgresRepository struct {
	Db *gorm.DB
}

func (userMysqlRepository *UserPostgresRepository) UpdateQuantityMovies(DNI int64, user model.User) (*model.User, error) {
	panic("implement me")
}

func (userMysqlRepository *UserPostgresRepository) GetByDNI(userDni int64) (model.User, error) {

	var userDb models.UserDb
	if userMysqlRepository.Db.First(&userDb, userDni).Error != nil {
		//exceptions.UserNotFound{ErrMessage: fmt.Sprintf("user with id=%d not found", userDni)}
		return model.User{}, nil
	}
	user := users_mapper.UserDbToUser(userDb)

	return user, nil
}

