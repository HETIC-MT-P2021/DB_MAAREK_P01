package controller

import (
	"github.com/JackMaarek/DS/repo"
	"github.com/JackMaarek/DS/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCustomerRecordById(c *gin.Context) {
	id := util.ParseStringToUint64(c.Param("id"))
	cm, err := repo.QueryCustomerById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "No content found")
	}

	c.JSON(http.StatusOK, cm)
}
