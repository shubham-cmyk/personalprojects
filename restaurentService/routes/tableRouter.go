package routes

import(
	controller "restaurentServiceProject/controllers"
	"github.com/gin-gonic/gin"

)

func TableRoutes(incomingRequest *gin.Engine){

	incomingRequest.GET("/tables",controller.GetTables())
	incomingRequest.GET("/tables/:table_id",controller.GetTable())
	incomingRequest.POST(/"tables",controller.CreateTable())
	incomingRequest.PATCH("/tables/:table_id",controller.UpdateTable())
	
}

