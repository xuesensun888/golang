package handlers

import (
	"operation/models"

	"github.com/gin-gonic/gin"
)

func Put_fault(c *gin.Context) {
	var fault models.FaultReport
	//获取动态路由传入的id
	id := c.Param("id")

	if err := models.Conn.DB.First(&fault, id).Error; err != nil {
		c.JSON(404, gin.H{
			"error": "user not found",
		})
	}

	//从请求体中绑定数据
	var input models.FaultReport

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	//更新字段值
	fault.Description = input.Description
	fault.Reporter = input.Reporter

	//保存更新后的数据
	if err := models.Conn.DB.Save(&fault).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "faild to update user",
		})
		return

	}

	c.JSON(200, fault)
}
