package main

import (
	"operation/conf"
	"operation/handlers"
	"operation/models"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置
	conf.LoadConfig("config.json")
	//初始化数据库链接
	models.Conn = models.NewConn()
	//定义一个默认的路由
	r := gin.Default()
	r.Use(cors.Default())
	apiGroup := r.Group("/api/v1/operation")
	{
		apiGroup.POST("/report_fault", handlers.Handler_fault)
		apiGroup.GET("/get_all_fault", handlers.GetAllFault)
		apiGroup.GET("/get_byid/:id", handlers.GetById)
		apiGroup.DELETE("/delete_id/:id", handlers.DeleteId)
		apiGroup.PUT("/put_id/:id", handlers.Put_fault)
		apiGroup.GET("/testapi", handlers.Hander_test)

	}
	r.Run(":8899")
}
