package process2

import (
	"encoding/json"
	"fmt"
	"gotest/chatroom/client/utils"
	"gotest/chatroom/common/message"
	"net"
)

type SmsProcess struct{}

//消息转发
func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//反序列化mes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json marshal err ====", err.Error())
		return
	}

	//序列化mes
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal mes", err.Error())
		return

	}
	//遍历useronlinx
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToOnlineUser(data []byte, conn net.Conn) {
	//发送消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败err=", err.Error())
	}
}
