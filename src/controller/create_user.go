package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/configuration/validation"
	"github.com/michaelcaxias/crud-golang/src/model/request"
	"github.com/michaelcaxias/crud-golang/src/model/response"
)

func CreateUser(context *gin.Context) {
	var userRequest request.UserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
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
