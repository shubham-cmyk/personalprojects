package routes

import (
	controller "restaurentServiceProject/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRequest *gin.Engine) {

	incomingRequest.GET("/orders", controller.GetOrders())
	incomingRequest.GET("/orders/:order_id", controller.GetOrder())
	incomingRequest.POST("/orders", controller.CreateOrder())
	incomingRequest.PATCH("/orders", controller.UpdateOrder())
}
