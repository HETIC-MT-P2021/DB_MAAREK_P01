package controller

import (
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/repo"
	"github.com/JackMaarek/DS/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrdersRecordByCustomer(c *gin.Context) {
	var o *[]models.PublicOrder
	var err error
	id := util.ParseStringToUint64(c.Param("id"))

	o, err = repo.QueryOrdersByCustomerId(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return
	}
	c.JSON(http.StatusOK, o)
}

func GetOrdersRecordById(c *gin.Context) {
	var o *models.Order
	var err error
	id := util.ParseStringToUint64(c.Param("id"))

	o, err = repo.QueryOrderById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return
	}
	c.JSON(http.StatusOK, o)
}

func GetOrdersDetailsRecordById(c *gin.Context) {
	var od *[]models.OrderDetails
	var err error
	id := util.ParseStringToUint64(c.Param("id"))

	od, err = repo.QueryOrderDetailsByOrderNumber(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return
	}

	c.JSON(http.StatusOK, od)
}
