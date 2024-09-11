package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
}

func main() {
	// 设置配置文件名（不带扩展名）
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置配置文件路径
	viper.AddConfigPath(".") // 当前目录

	// 自动绑定环境变量
	viper.AutomaticEnv()

	// 使用替换器来处理环境变量名中的下划线
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 设置默认值
	viper.SetDefault("ssh_agent_pid", 8080)
	viper.SetDefault("server.host", "localhost")

	// 打印环境变量以进行调试
	fmt.Println("Environment Variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	// 读取配置文件（如果存在）
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, using default values and environment variables")
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	// 打印 Viper 中的配置值以进行调试
	fmt.Println("Viper Configuration:")
	fmt.Printf("server.port: %d\n", viper.GetInt("ssh_agent_pid"))
	fmt.Printf("server.host: %s\n", viper.GetString("server.host"))

}

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// // 创建websocket升级器
// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// //websocket处理函数

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	//升级http连接到websocket
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer conn.Close()
// 	//读取并回显消息
// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			break

// 		}
// 		log.Printf("recevied %s", message)
// 		if err := conn.WriteMessage(messageType, message); err != nil {
// 			log.Println(err)
// 			break
// 		}
// 	}

// }
// func main() {
// 	//注册默认路由及绑定函数
// 	http.HandleFunc("/ws", wsHandler)
// 	log.Println("server started on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }
