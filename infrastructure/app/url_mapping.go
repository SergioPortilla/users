package app

import (
	"github.com/ceiba-meli-demo/users/infrastructure/controllers"
)

func mapUrls(handler controllers.RedirectUserHandler) {

	router.GET("/users/:user_id", handler.Get)
	router.PUT("/users/:user_id", handler.Update)

}