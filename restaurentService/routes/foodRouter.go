package routes

import (
	controller "restaurentServiceProject/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRequest *gin.Engine) {

	incomingRequest.GET("/foods", controller.GetFoods())
	incomingRequest.GET("/foods/:food_id", controller.GetFood())
	incomingRequest.POST("/foods", controller.CreateFood())
	incomingRequest.PATCH("/foods/:food_id", controller.UpdateFood())

}
