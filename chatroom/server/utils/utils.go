package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gotest/chatroom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的数据")
	//conn.read 在conn没有关闭情况下才会阻塞
	//如何客户端关闭了conn就不会阻塞

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	//根据buf[:4] 转成unit32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[:4])
	//根据pkglen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	//把pkglen反序列化成>>>>message.message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.unmarchal err", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度给对方
	pkglen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkglen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write(bytes) fail", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.write(bytes) fail ", err)
		return
	}
	return
}
