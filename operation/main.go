package main

import (
	"operation/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//设置上传文件目录的路径
	//r.Static("/uploads", "./uploads")
	routers.AdminRoutersInit(r)
	routers.UserRoutesInit(r)
	r.Run(":8989")

}
