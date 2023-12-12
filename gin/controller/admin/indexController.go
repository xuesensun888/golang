package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	//
	username, _ := c.Get("username")
	fmt.Println(username)
	//username是一个空接口类型，顾要使用断言类型转换
	v, ok := username.(string)
	if ok != true {
		c.String(200, "后台首页-获取用户名失败")
	} else {
		c.String(200, "后台首页，用户名"+v)
	}
}
