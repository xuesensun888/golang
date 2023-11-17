package main

import (
	"fmt"
	"gotest/chatroom/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	//循环客户端发送的消息
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端协成错误=err", err)
		return
	}
}
func init() {
	initPool("172.24.0.21:6379", 16, 0, 300*time.Second)
	initUserDao()
}

//这里我们编写一个函数，完成对userdao的初始化任务
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	fmt.Println("服务器正在监听端口9000")
	listen, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("listen error", err)
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn error", err)
		}
		go process(conn)
	}
}
