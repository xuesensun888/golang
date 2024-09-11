package handlers

import (
	"net/http"
	"operation/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetById(c *gin.Context) {
	var fault models.FaultReport
	id := c.Param("id")
	//根据id获取特定用户
	if err := models.Conn.DB.First(&fault, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
		}
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": fault,
	})
}
