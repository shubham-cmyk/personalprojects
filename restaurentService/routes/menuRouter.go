package routes

import (
	controller "restaurentServiceProject/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRequest *gin.Engine) {

	incomingRequest.GET("/menus", controller.GetMenus())
	incomingRequest.GET("/menus/:menu_id", controller.GetMenu())
	incomingRequest.POST("/menus", controller.CreateMenu())
	incomingRequest.PATCH("/menus/:menu_id", controller.UpdateMenu())

}
