package controller

import (
	"fmt"
	"operation/models"

	"github.com/gin-gonic/gin"
)

type FaultController struct {
}

func (FaultController) Fault(c *gin.Context) {
	var fault models.Fault
	err := c.ShouldBind(&fault)
	if err != nil {
		fmt.Println("结构体绑定报错", err)
	}

	models.DB.Create(fault)

	fmt.Println(fault)
	c.String(200, "增加数据成功")

}
func (FaultController) GetFault(c *gin.Context) {
	//查找全部
	userList := []models.Fault{}
	id := c.Param("id")
	//查询所有用户的数据，把结果保存到userlist中
	models.DB.Where("id=?", id).Find(&userList)
	c.JSON(200, userList)
}
