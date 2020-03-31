package controllers

import (
	"github.com/ceiba-meli-demo/users/application/usescases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	GetUserUseCase          usescases.GetByDniUseCase
	UseCaseUpdateUser       usescases.UpdateUserUseCase
}
func (h *Handler) Get(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		return
	}
	user, errGet := h.GetUserUseCase.Handler(userId)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}

	c.JSON(http.StatusOK, user)

}
