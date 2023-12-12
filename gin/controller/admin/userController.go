package admin

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "用户列表1")
}
func (con UserController) Add(c *gin.Context) {
	c.String(200, "用户添加2")
}
func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户修改3")
}

