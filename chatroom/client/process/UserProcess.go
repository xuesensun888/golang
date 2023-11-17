package process

import (
	"encoding/json"
	"fmt"
	"gotest/chatroom/common/message"
	"net"
	"os"

	"gotest/chatroom/client/utils"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//开始定协议
	conn, err := net.Dial("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("net.dial error ====", err)
		return
	}
	//延迟关闭
	defer conn.Close()
	//2 准备通过conn发送消息给服务

	var mes message.Message
	mes.Type = message.LoginMesType
	//3 创建一个loginmes结构体
	var LoginMes message.LoginMes
	LoginMes.UserId = userId
	LoginMes.UserPwd = userPwd
	//4 将loginmes序列化
	data, err := json.Marshal(LoginMes)
	if err != nil {
		fmt.Println("loginmes json.marshal err=", err)
		return
	}
	//5 将data切片赋值给mes.Data
	mes.Data = string(data)
	fmt.Println("mes.data ====", mes.Data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.marchal err=", err)
		return
	}
	//7 到这个时候data就是我们要发送的数据了
	//7.1 首先先把data长度发送给服务器
	//先获取data的长度》〉》转成一个表示长度byte切片
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("write pkg is error", err)
		return
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readpkg is error", err)
		return
	}
	//将mes的data部分反序列化成功loginresmes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			//如果要求不显示自己在下边确认
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
		}
		//fmt.Println("登录成功")
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 || loginResMes.Code == 403 {
		fmt.Println(loginResMes.Error)
	}
	return

}
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//连接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("net.dial error", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType
	//定义结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//序列化registermes
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.marshal err", err)
	}
	//将data赋值给mes.data
	mes.Data = string(data)
	//mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marchal mes fail", err)
		return
	}
	//这时data就是我们要发送的数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送消息错误err=", err)
		return
	}
	mes, err = tf.ReadPkg()
	//将收到的package 反序列化成registerResmes
	//定义RegisterResMes结构题
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json unmarchal err", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}
