package handlers

import (
	"operation/models"

	"github.com/gin-gonic/gin"
)

func DeleteId(c *gin.Context) {
	var fault models.FaultReport
	id := c.Param("id")
	//根据id获取特定用户
	if err := models.Conn.DB.First(&fault, id).Error; err != nil {
		c.JSON(404, gin.H{
			"Error": "user not found",
		})
		return
	}
	//删除用户
	if err := models.Conn.DB.Delete(&fault).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "faild to delete user",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user deleted susscesful",
	})
}
