package app

import (
	"github.com/ceiba-meli-demo/users/application/usescases"
	"github.com/ceiba-meli-demo/users/domain/ports"
	"github.com/ceiba-meli-demo/users/infrastructure/adapters/repository/users"
	"github.com/ceiba-meli-demo/users/infrastructure/app/middlewares/error_handler"
	"github.com/ceiba-meli-demo/users/infrastructure/controllers"
	"github.com/ceiba-meli-demo/users/infrastructure/database_client"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(error_handler.ErrorHandler())
	userRepository := getUsersRepository()
	var handler = createHandler(userRepository)
	mapUrls(handler)
	getUsersRepository()
	router.Run(":8086")
}

func createHandler(userRepository ports.UsersRepository) controllers.RedirectUserHandler {
	return newHandler(newGetUserUseCase(userRepository), newUpdateUserUseCase(userRepository))
}

func newGetUserUseCase(repository ports.UsersRepository) usescases.GetByDniUseCase {
	return &usescases.UseCaseGetByDni{
		UserRepository: repository,
	}
}

func newUpdateUserUseCase(repository ports.UsersRepository) usescases.UpdateUserUseCase {
	return &usescases.UseCaseUpdateUser{
		UserRepository: repository,
	}
}

func newHandler(getUserUseCase usescases.GetByDniUseCase, updateUserUseCase usescases.UpdateUserUseCase) controllers.RedirectUserHandler {
	return &controllers.Handler{
		GetUserUseCase:    getUserUseCase,
		UseCaseUpdateUser: updateUserUseCase,
	}
}

func getUsersRepository() ports.UsersRepository {
	return &users.UserPostgresRepository{
		Db: database_client.GetDatabaseInstance(),
	}
}
