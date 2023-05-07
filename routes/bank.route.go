package routes

import (
	"github.com/celpostgress-api/controller"

	"github.com/gin-gonic/gin"
)

func BankRouter(bankController controller.IBankController, group *gin.RouterGroup) *gin.RouterGroup {

	group.GET("/", bankController.Find)
	group.POST("/q", bankController.Find)
	group.GET("/:oid", bankController.FindById)
	group.POST("/create", bankController.Create)
	group.PUT("/:oid", bankController.Update)
	group.DELETE("/delete/:oid", bankController.Delete)

	return group
}
