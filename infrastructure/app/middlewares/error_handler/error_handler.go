package error_handler

import (
	"github.com/ceiba-meli-demo/users/domain/exceptions"
	"github.com/ceiba-meli-demo/users/infrastructure/utils/logger"
	"github.com/ceiba-meli-demo/users/infrastructure/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		// Use reflect.TypeOf(err.Err) to known the type of your error
		if _, ok := errors.Cause(err.Err).(exceptions.Validator); ok {
			restErr := rest_errors.NewBadRequestError(err.Error())
			logger.Error(restErr.Message(), restErr)
			c.JSON(restErr.Status(), restErr)
			return
		}
		// Use reflect.TypeOf(err.Err) to known the type of your error
		if _, ok := errors.Cause(err.Err).(exceptions.UserNotFound); ok {
			restErr := rest_errors.NewNotFoundError(err.Error())
			logger.Error(restErr.Message(), restErr)
			c.JSON(restErr.Status(), restErr)
			return
		}
		restErr := rest_errors.NewInternalServerError(err.Error(), err)
		logger.Error(restErr.Message(), restErr)
		c.JSON(restErr.Status(), restErr)
	}
}
