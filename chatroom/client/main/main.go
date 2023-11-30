package main

import (
	"fmt"
	"gotest/chatroom/client/process"
	"os"
)

// 定义变量
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	for {
		fmt.Println("--------欢迎登录多人聊天系统------------")
		fmt.Println("\t\t\t 1.登录聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3): ")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			up := process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("请输入注册用户的id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入注册用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的昵称:")
			fmt.Scanf("%s\n", &userName)
			//调用userprocess 完成注册的请求
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你输入的有错误")
		}
	}
}
