package main

import (
	"remotehttp/conf"
	"remotehttp/handler"
	"remotehttp/model"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.LoadConfig("config.json")

	model.Conn = model.NewConn()

	r := gin.Default()
	v := r.Group("/api/v1/io/app")
	{
		v.GET("/:app", handler.Version)
		v.HEAD("/download/:name", handler.Header)
		//	下载包
		v.GET("/download/:name", handler.Download)

	}
	r.Run(":9092")
}
