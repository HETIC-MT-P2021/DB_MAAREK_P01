package controller

import (
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/repo"
	"github.com/JackMaarek/DS/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEmployeeRecordById(c *gin.Context) {
	var err error
	var e *models.Employee
	id := util.ParseStringToUint64(c.Param("id"))

	e, err = repo.QueryEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return
	}

	c.JSON(http.StatusOK, e)
}
