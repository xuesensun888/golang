package main

import (
	"fmt"
	"time"
)

// type Article struct {
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

//	func UnixToTime(timestamp int) string {
//		fmt.Println(timestamp)
//		t := time.Unix(int64(timestamp), 0)
//		return t.Format("2006-01-02 15:04:05")
//	}
//
//	func Println(str1 string, str2 string) string {
//		return str1 + str2
//	}

func main() {
	// //创建默认路由引擎
	// r := gin.Default()
	// //加载templates中所有模版文件，使用不同目录下名称相同的模版， 一定要放在配置路由钱
	// r.LoadHTMLGlob("templates/**/*")
	// //配置静态web目录，第一个参数表示路由。第二个参数表示映射目录
	// r.Static("/static", "./static")
	// //全局中间件
	// //r.Use(initMiddlewareTwo, initMiddleware)
	fmt.Println(time.Month.String(12))
	// routers.AdminRoutersInit(r)
	// routers.ApiRoutersInit(r)
	// routers.DefaultRoutersInit(r)
	// r.Run()
}
