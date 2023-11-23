package model

import (
	"gotest/chatroom/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
