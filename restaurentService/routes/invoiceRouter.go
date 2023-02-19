package routes

import (
	controller "restaurentServiceProject/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRequest *gin.Engine) {
	incomingRequest.GET("/invoices", controller.GetInvoices())
	incomingRequest.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRequest.POST("/invoices", controller.CreateInvoice())
	incomingRequest.PATCH("invoices/:invoice_id", controller.UpdateInvoice())

}
