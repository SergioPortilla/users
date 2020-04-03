package users

import (
	"errors"
	"fmt"
	"github.com/ceiba-meli-demo/users/domain/model"
	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/models"
	"github.com/ceiba-meli-demo/users/infrastructure/mappers/users_mapper"
	"github.com/jinzhu/gorm"
)

type UserSqlRepository struct {
	Db *gorm.DB
}

func (userPostgresRepository *UserSqlRepository) UpdateQuantityMovies(DNI int64) (*model.Message, error) {

	userDb, error := findByDNI(userPostgresRepository, DNI)
	if error != nil {
		return nil, error
	}

	if userPostgresRepository.Db.Model(&userDb).Update("quantity_movies", gorm.Expr("quantity_movies  + ?", 1)).Error != nil {
		return nil, errors.New(fmt.Sprintf("error when updated user %v", DNI))
	}
	return &model.Message{
		Msg:  "quantity movies update",
		Done: true,
	}, nil

}

func (userPostgresRepository *UserSqlRepository) GetByDNI(DNI int64) (*model.User, error) {

	userDb, error := findByDNI(userPostgresRepository, DNI)
	if error != nil {
		return nil, error
	}
	user := users_mapper.UserDbToUser(*userDb)
	return &user, nil

}

func findByDNI(userPostgresRepository *UserSqlRepository, DNI int64) (*models.UserDb, error) {

	var userDb models.UserDb
	if userPostgresRepository.Db.First(&userDb, DNI).Error != nil {
		return nil, errors.New(fmt.Sprintf("user not found %v", DNI))
	} else {
		return &userDb, nil
	}

}
