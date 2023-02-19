package routes

import (
	controller "restaurentServiceProject/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRequest *gin.Engine) {

	incomingRequest.GET("/orderItems", controller.GetOrderItems())
	incomingRequest.GET("/orderItems/:orderItem_id", controller.GetOrderItems())
	incomingRequest.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRequest.POST("/orderItems", controller.CreateOrderItem())
	incomingRequest.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
