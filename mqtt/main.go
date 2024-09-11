// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// 	"github.com/shirou/gopsutil/cpu"
// 	"github.com/shirou/gopsutil/disk"
// 	"github.com/shirou/gopsutil/mem"
// )

// // 定义系统统计数据的结构体
// type SystemStats struct {
// 	CPU    float64 `json:"cpu"`
// 	Memory float64 `json:"memory"`
// 	Disk   float64 `json:"disk"`
// }

// // 获取系统统计数据
// func getSystemStats() (SystemStats, error) {
// 	var stats SystemStats

// 	// 获取 CPU 使用率
// 	cpuPercents, err := cpu.Percent(time.Second, false)
// 	if err != nil {
// 		return stats, err
// 	}
// 	stats.CPU = cpuPercents[0]

// 	// 获取内存使用情况
// 	v, err := mem.VirtualMemory()
// 	if err != nil {
// 		return stats, err
// 	}
// 	stats.Memory = v.UsedPercent

// 	// 获取硬盘使用情况
// 	d, err := disk.Usage("/")
// 	if err != nil {
// 		return stats, err
// 	}
// 	stats.Disk = d.UsedPercent

// 	return stats, nil
// }

// // 消息处理函数
// var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
// 	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
// }

// // 连接处理函数
// var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
// 	fmt.Println("Connected")
// }

// // 连接丢失处理函数
// var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
// 	fmt.Printf("Connect lost: %v", err)
// }

// func main() {
// 	// MQTT 服务器选项
// 	opts := mqtt.NewClientOptions()
// 	opts.AddBroker(":1883") // 这里可以换成你自己的 MQTT 服务器地址
// 	opts.SetClientID("mqttx_9e681512")
// 	// opts.SetUsername("your-username") // 可选
// 	// opts.SetPassword("your-password") // 可选
// 	opts.OnConnect = connectHandler
// 	opts.OnConnectionLost = connectLostHandler

// 	// 创建客户端
// 	client := mqtt.NewClient(opts)
// 	if token := client.Connect(); token.Wait() && token.Error() != nil {
// 		fmt.Println(token.Error())
// 		os.Exit(1)
// 	}

// 	// 订阅主题
// 	topic := "sxs/topic"
// 	token := client.Subscribe(topic, 1, messagePubHandler)
// 	token.Wait()
// 	fmt.Printf("Subscribed to topic: %s\n", topic)

// 	// 循环获取系统数据并发布到 MQTT
// 	for {
// 		stats, err := getSystemStats()
// 		if err != nil {
// 			log.Printf("Error getting system stats: %v", err)
// 			continue
// 		}

// 		msg, err := json.Marshal(stats)
// 		if err != nil {
// 			log.Printf("Error marshaling message: %v", err)
// 			continue
// 		}

// 		token := client.Publish(topic, 0, false, msg)
// 		token.Wait()

// 		log.Printf("Published message: %s", msg)

// 		time.Sleep(10 * time.Second)
// 	}

// 	// 等待消息接收
// 	time.Sleep(6 * time.Second)

//		// 断开连接
//		client.Disconnect(250)
//	}
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许来自任何来源的请求
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr())
	fmt.Println(conn.RemoteAddr())
	fmt.Println("Client connected")

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error while reading message:", err)
			break
		}

		fmt.Printf("Received: %s\n", msg)

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Error while writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	serverAddr := "localhost:8099"
	fmt.Printf("WebSocket server started at ws://%s\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
