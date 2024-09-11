package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://127.0.0.1:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial error:", err)

	}
	defer conn.Close()
	fmt.Println("连接成功")
	//发送消息
	message := []byte("hello, websocket")
	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		log.Fatal("write error:", err)
	}
	//读取服务器的响应
	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("read error:", err)

	}

	log.Printf("received %s", response)
	//等待一段时间后关闭连接
	time.Sleep(time.Second)
}
