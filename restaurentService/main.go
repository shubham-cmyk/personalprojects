package main

import (
	"restaurentServiceProject/middleware"
	"restaurentServiceProject/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var database *gorm.DB

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":8000")

}
