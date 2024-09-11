package main

import (
	"fmt"
	"proto/userService"

	"google.golang.org/protobuf/proto"
)

func main() {

	u := &userService.Userinfo{
		Username: "张三",
		Age:      20,
		Hobby:    []string{"吃饭", "睡觉", "打球"},
	}
	data, _ := proto.Marshal(u)
	fmt.Println(u.Age)
	fmt.Println(u.Hobby)
	//protobuf 序列化
	fmt.Println(data)
	//protobuf 反序列化
	var user userService.Userinfo
	proto.Unmarshal(data, &user)
	fmt.Println(user)
}
