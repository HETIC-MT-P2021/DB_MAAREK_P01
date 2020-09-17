package router

import (
	"github.com/JackMaarek/DS/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/customers", controller.GetCustomerRecord)
	router.GET("/customers/:id", controller.GetCustomerRecordById)
}
