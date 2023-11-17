package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gotest/chatroom/common/message"
	"net"
)

//这里将方法关联到结构体上
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("客户端正在读取服务端发送的数据")
	//conn.read 在conn没有关闭情况下 才会阻塞
	//如果客户端关闭conn则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	//根据buf[:4]转成uint32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])
	//根据pkglen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	//把pkglen序列化成》〉》〉message.message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json unmachal err= ", err)
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度

	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(bytes) fail", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write(byte) fail", err)
		return
	}
	return
}
