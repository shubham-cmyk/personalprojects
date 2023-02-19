package routes

import (
	controller "restaurentServiceProject/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRequest *gin.Engine) {

	incomingRequest.GET("/users", controller.GetUsers())
	incomingRequest.GET("/users/:user_id", controller.GetUser())
	incomingRequest.POST("/users/signup", controller.SignUp())
	incomingRequest.POST("users/login", controller.Login())

}
