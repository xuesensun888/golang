package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认路由
	r := gin.Default()
	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "你好gin")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})
	r.Run()
}
