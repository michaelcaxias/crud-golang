package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/configuration/restError"
)

func FindUserByID(context *gin.Context) {
	err := restError.ThrowNotFoundError("Usuário não encontrado")

	context.JSON(err.Status, err)
}

func FindAllUsers(context *gin.Context) {
	err := restError.ThrowNotFoundError("Usuários não encontrados")

	context.JSON(err.Status, err)
}
