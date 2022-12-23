package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/configuration/logger"
	"github.com/michaelcaxias/crud-golang/src/configuration/validation"
	"github.com/michaelcaxias/crud-golang/src/model/request"
	"github.com/michaelcaxias/crud-golang/src/model/response"
	"go.uber.org/zap"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	logger.Info("Init CreateUser controler", 
	zap.String("journey", "createUser"),
)

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restError := validation.ValidateUserError(err)

		context.JSON(restError.Status, restError)
		return
	}

	response := response.UserResponse{
		ID: "1",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	context.JSON(http.StatusCreated, response)
}
