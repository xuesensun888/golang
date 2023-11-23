package process

import (
	"encoding/json"
	"fmt"
	"gotest/chatroom/client/utils"
	"gotest/chatroom/common/message"
)

type SmsProcess struct{}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//定义一个mes结构体实例
	var mes message.Message
	mes.Type = message.SmsMesType

	//定义一个sms结构体实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = curUser.UserId
	smsMes.UserStatus = curUser.UserStatus
	//序列化smsmes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json marshal err =", err.Error())
		return
	}
	//将data赋值给mes
	mes.Data = string(data)

	//序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marchal mes err =", err.Error())
		return

	}
	//将消息发送出去
	tf := utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sendgroupmes", err.Error())
		return
	}
	return
}
