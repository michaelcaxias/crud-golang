package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/configuration/rest_error"
)

func FindUserByID(context *gin.Context) {
	err := rest_error.ThrowNotFoundError("User not found")

	context.JSON(err.Status, err)
}

func FindAllUsers(context *gin.Context) {
	err := rest_error.ThrowNotFoundError("Users not found")

	context.JSON(err.Status, err)
}
