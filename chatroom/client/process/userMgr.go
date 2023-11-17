package process

import (
	"fmt"
	"gotest/chatroom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

//编写一个方法处理notifyuserstatus

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
func outputOnlineUser() {
	fmt.Println("当前在线列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户的id:\t", id)
	}
}
