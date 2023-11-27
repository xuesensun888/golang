package process

import (
	"encoding/json"
	"fmt"
	"gotest/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) {
	//反序列化mes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.umarchal err", err.Error())
		return
	}
	//显示信息
	fmt.Printf("用户id:\t%d 对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println()
}
