package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients    = make(map[*websocket.Conn]bool)
	clientsMux sync.Mutex
)

// type PackageMessage struct {
// 	FileName string `json:"file_name"`
// 	Content  []byte `json:"content"`
// }

func handleConnect(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error while upgrading connection: ", err)
	}
	defer conn.Close()

	//添加互斥锁写入map数据,以免发生资源竞争
	clientsMux.Lock()
	clients[conn] = true //将新的连接添加到clients map中
	clientsMux.Unlock()
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("error while read message", nil)
			clientsMux.Lock()
			delete(clients, conn)
			clientsMux.Unlock()
			break
		}
		//广播消息到所有的其他链接
		clientsMux.Lock()
		for client := range clients {
			if client != conn {
				err := client.WriteMessage(messageType, msg)
				if err != nil {
					log.Println("error while send message to client:", client)
					client.Close()
					delete(clients, client)
				}
			}

		}
		clientsMux.Unlock()
	}

	//定义包名
	// packageFileName := "cargo-robot-1.1.5.deb"
	// file, err := os.Open(packageFileName)
	// if err != nil {
	// 	log.Fatal("Error opening .deb file: ", err)
	// }
	// defer file.Close()

	// // Read file content
	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	log.Fatal("Error getting file info: ", err)
	// }
	// fileSize := fileInfo.Size()
	// buffer := make([]byte, fileSize)

	// _, err = file.Read(buffer)
	// if err != nil && err != io.EOF {
	// 	log.Fatal("Error reading file: ", err)
	// }
	// // Send package name and content
	// msg := PackageMessage{
	// 	FileName: packageFileName,
	// 	Content:  buffer,
	// }
	// err = conn.WriteJSON(msg)
	// if err != nil {
	// 	log.Fatal("Error sending message: ", err)
	// }

	// Notify client that the file transmission is complete
	// conn.WriteMessage(websocket.TextMessage, []byte("EOF"))

}

func main() {
	http.HandleFunc("/ws", handleConnect)
	serverAddr := ":8080"
	fmt.Printf("Server started at %s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
