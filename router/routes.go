package router

import (
	"github.com/JackMaarek/DS/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	apiV1 := router.Group("/v1")

	// -------- Customers -----------
	customers := apiV1.Group("/customers")

	customers.GET("/:id", controller.GetCustomerRecordById)
	customers.GET("/:id/orders", controller.GetOrdersRecordByCustomer)
	// -------- End Customers -----------

	// -------- Orders -----------
	orders := apiV1.Group("/orders")

	orders.GET("/:id", controller.GetOrdersRecordById)
	orders.GET("/:id/details", controller.GetOrdersDetailsRecordById)
	// -------- End Orders -----------

	// -------- Employees -----------
	employees := apiV1.Group("/employees")

	employees.GET("/:id", controller.GetEmployeeRecordById)
	// -------- End Employees -----------

	// -------- Offices -----------
	offices := apiV1.Group("/offices")

	offices.GET("/:id", controller.GetOfficeRecordById)
	// -------- End Offices -----------

}
