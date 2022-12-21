package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/configuration/restError"
)

func DeleteUser(context *gin.Context) {
	err := restError.ThrowNotFoundError("Usuário não encontrado")

	context.JSON(err.Status, err)
}
