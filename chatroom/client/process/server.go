package process

import (
	"encoding/json"
	"fmt"
	"gotest/chatroom/client/utils"
	"gotest/chatroom/common/message"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("--------恭喜登录成功---------")
	fmt.Println("--------1.显示在线用户------------")
	fmt.Println("--------2.发送消息------------")
	fmt.Println("--------3.信息列表------------")
	fmt.Println("--------4.退出系统------------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("发送消息")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择了退出系统。。")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}
}

//和服务器保持通讯
func ServerProcessMes(conn net.Conn) {
	//创建一个transfer实例 不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("ts.readpkg err=", err)
			return
		}
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//1.取出notifyuserstatusmes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2 把这个用户信息保存到客户map中
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了位置的消息类型")
		}

	}
}
