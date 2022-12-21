package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/configuration/rest_error"
	"github.com/michaelcaxias/crud-golang/src/model/request"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		restError := rest_error.ThrowBadRequestError(
			fmt.Sprintf("Existem alguns campos inválidos, error %s", err.Error()),
		)

		context.JSON(restError.Status, restError)
		return
	}
}
