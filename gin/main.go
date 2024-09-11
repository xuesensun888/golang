package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name") // 获取 POST 请求中的名字字段值

		// 在后端处理获取到的数据
		println("接收到的名字是:", name)

		// 可以进行其他处理或者返回响应给前端
		
		// c.String(http.StatusOK, "成功接收到名字: %s", name)
	})

	router.Run(":8080") // 启动服务器，监听本地的 8080 端口
}
