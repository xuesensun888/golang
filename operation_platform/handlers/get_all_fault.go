package handlers

import (
	"net/http"
	"operation/models"

	"github.com/gin-gonic/gin"
)

func GetAllFault(c *gin.Context) {
	var faults []models.FaultReport
	err := models.Conn.DB.Find(&faults).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//返回查询结果
	c.JSON(http.StatusOK, gin.H{
		"data": faults,
	})
}
