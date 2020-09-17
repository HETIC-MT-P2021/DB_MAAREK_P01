package controller

import (
	"github.com/JackMaarek/DS/repo"
	"github.com/JackMaarek/DS/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomerRecord(c *gin.Context) {
	repo.QueryAllCustomers()

	c.JSON(http.StatusOK, "Shalom")
}

func GetCustomerRecordById(c *gin.Context) {
	id := util.ConvertStringToInt(c.Param("id"))
	cm := repo.QueryCustomerById(id)

	c.JSON(http.StatusOK, cm)
}
