package app

import (
	"github.com/ceiba-meli-demo/users/domain/ports"
	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/users"
	"github.com/ceiba-meli-demo/users/infrastructure/database_client"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	getUsersRepository()
	router.Run(":8086")
}

func getUsersRepository() ports.UsersRepository {
	return &users.UserPostgresRepository{
		Db: database_client.GetDatabaseInstance(),
	}
}