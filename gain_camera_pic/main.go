package main

import (
	"camera/src/handler"
	"camera/src/log"
	"camera/src/modules"
	"fmt"
	"os"
)

func main() {
	//conf.LoadConfig("config.json")

	var key int
	for {
		fmt.Println("---------欢迎使用自动化工具-----------")
		fmt.Println("\t 1. topic获取")
		fmt.Println("\t 2. 下载camera JPG")
		fmt.Println("\t 3. 生成camera JPG")
		fmt.Scan(&key)
		switch key {
		case 1:
			log.Logger.Println("topic正在获取")
			handler.Getopic()
		case 2:
			log.Logger.Println("正在下载camera JPG")
		case 3:
			log.Logger.Println("正在生成canmera JPG")
			modules.Cam_Generate()

		default:
			log.Logger.Println("你输入的有误")
			os.Exit(11)
		}

	}
}
