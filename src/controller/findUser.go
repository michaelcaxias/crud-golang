package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/configuration/restError"
)

func FindUserByID(context *gin.Context) {
	err := restError.NewNotFoundError("Usuário não encontrado")

	context.JSON(err.Status, err)
}

func FindAllUsers(context *gin.Context) {
	err := restError.NewNotFoundError("Usuários não encontrados")

	context.JSON(err.Status, err)
}
