package controllers

import (
	"net/http"
	"strconv"

	"github.com/ceiba-meli-demo/users/application/usescases"
	"github.com/ceiba-meli-demo/users/infrastructure/utils/rest_errors"
	"github.com/gin-gonic/gin"
)

type RedirectUserHandler interface {
	Get(c *gin.Context)
	Update(c *gin.Context)
}

type Handler struct {
	GetUserUseCase    usescases.GetByDniUseCase
	UseCaseUpdateUser usescases.UpdateUserUseCase
}

func (h *Handler) Update(c *gin.Context) {
	dni, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := rest_errors.NewBadRequestError("user_id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	user, errGet := h.UseCaseUpdateUser.Handler(dni)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) Get(c *gin.Context) {

	dni, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := rest_errors.NewBadRequestError("user_id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	user, errGet := h.GetUserUseCase.Handler(dni)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}

	c.JSON(http.StatusOK, user)

}
