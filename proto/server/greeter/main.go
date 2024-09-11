package main

import (
	"context"
	"fmt"
	"net"
	"proto/greeter"

	"google.golang.org/grpc"
)

type Hello struct{}

func (con *Hello) SayHello(c context.Context, req *greeter.HelloReq) (*greeter.HelloRes, error) {
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "你好" + req.Name,
	}, nil
}
func main() {
	//初始化grpc
	grpcServer := grpc.NewServer()
	//2 注册服务
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	//3 设置监听 指定ip port
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	//启动服务
	grpcServer.Serve(listener)
}
