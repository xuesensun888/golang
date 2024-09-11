package main

import (
	"fmt"
	"net/rpc"
)

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

func main() {
	//用rpc.dial 和rpc微服务端建立连接
	conn, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
	}
	//当客户端退出时关闭连接
	defer conn.Close()
	//调用远程函数
	var reply AddGoodsRes

	err2 := conn.Call("goods.AddGoods", AddGoodsReq{
		Id:      10,
		Title:   "商品标题",
		Price:   23.6,
		Content: "商品详情",
	}, &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	//获取微服务返回的数据
	fmt.Printf("%v\n", reply)

	var goodsDate GetGoodsRes

	err3 := conn.Call("goods.GetGoods", GetGoodsReq{
		Id: 12,
	}, &goodsDate)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Printf("%v\n", goodsDate)
}
