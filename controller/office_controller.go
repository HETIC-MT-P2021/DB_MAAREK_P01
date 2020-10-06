package controller

import (
	"github.com/JackMaarek/DS/models"
	"github.com/JackMaarek/DS/repo"
	"github.com/JackMaarek/DS/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOfficeRecordById(c *gin.Context) {
	var err error
	var o *models.Office
	id := util.ParseStringToUint64(c.Param("id"))

	o, err = repo.QueryOfficeByID(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "")
		return
	}

	c.JSON(http.StatusOK, o)
}
