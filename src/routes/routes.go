package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelcaxias/crud-golang/src/controller"
)

func InitRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/user/:id", controller.FindUserByID)
	routerGroup.GET("/user/", controller.FindAllUsers)
	routerGroup.POST("/user/", controller.CreateUser)
	routerGroup.PUT("/user/:id", controller.UpdateUser)
	routerGroup.DELETE("/user/:id", controller.DeleteUser)
}
