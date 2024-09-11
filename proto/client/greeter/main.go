package main

import (
	"context"
	"fmt"
	"log"
	"proto/greeter"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//客户端连接
	clinetconn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接grpc服务器失败", err)
	}
	defer clinetconn.Close()

	//创建rpc客户端
	client := greeter.NewGreeterClient(clinetconn)
	res, err := client.SayHello(context.Background(), &greeter.HelloReq{
		Name: "",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", res)
}
