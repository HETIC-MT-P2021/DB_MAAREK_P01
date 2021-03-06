package controller

import (
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/repo"
	"github.com/JackMaarek/DS/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomerRecordById(c *gin.Context) {
	var err error
	var cm *models.Customer
	id := util.ParseStringToUint64(c.Param("id"))

	cm, err = repo.QueryCustomerById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return
	}
	c.JSON(http.StatusOK, cm)
}
