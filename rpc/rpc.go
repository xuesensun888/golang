package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义一个远程调用的方法
// type Hello struct{}

// 1 方法只能有两个序列化的参数，其中第二个参数是指针类型
// req 表示获取客户端传过来的数据
// res 表示给客户端返回的数据
// 2 方法要返回error类型同时必须是公开的方法
// 3req 和res类型不能是channel 、func函数均不能进行序列话
//
//	func (this Hello) SayHello(req string, res *string) error {
//		*res = "你好" + req
//		return nil
//	}
type Goods struct{}
type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}
type AddGoodsRes struct {
	Success bool
	Message string
}
type GetGoodsReq struct {
	Id int
}
type GetGoodsRes struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

func (con Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	fmt.Println(req)
	*res = AddGoodsRes{
		Success: true,
		Message: "增加数据成功",
	}
	return nil
}
func (con Goods) GetGoods(req GetGoodsReq, res *GetGoodsRes) error {
	fmt.Printf("%#v", req)
	*res = GetGoodsRes{
		Id:      12,
		Title:   "服务器获取的数据",
		Price:   24.5,
		Content: "我是服务器数据库获取的内容",
	}
	return nil
}
func tcp() {
	//创建rpc服务
	rpcServer := rpc.NewServer()
	//1 注册rpc服务
	err1 := rpcServer.RegisterName("goods", new(Goods))
	if err1 != nil {
		fmt.Println(err1)
	}
	//2 监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8090")
	if err2 != nil {
		fmt.Println(err2)
	}
	//3 应用退出的时候关闭监听端口
	defer listener.Close()
	fmt.Println("等待建立连接")
	rpcServer.Accept(listener)

}
