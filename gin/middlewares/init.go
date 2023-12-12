package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	//设置值 和对应控制器之间共享数据
	c.Set("username", "张三")
	//gin中间件中使用gorouting
	//当在中间件或handler中启动新的gorouting时，不能使用原始的上下文（c *gin.Context）
	//必须使用只读副本C.COPY
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("gin中使用gorouting" + cCp.Request.URL.Path)
	}()
}
