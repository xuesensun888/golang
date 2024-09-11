// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// //创建一个通道，用于接受信号
// 	// sigs := make(chan os.Signal, 1)
// 	// //注册通道以接收特定的信号
// 	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
// 	// //启动一个gorountine等待信号
// 	// go func() {
// 	// 	sig := <-sigs
// 	// 	fmt.Println()
// 	// 	fmt.Println("received signal:", sig)
// 	// 	os.Exit(0)
// 	// }()
// 	// //主gorountine等待
// 	// fmt.Println("press ctrl+c to exit")
// 	// for {
// 	// 	fmt.Print(".")
// 	// 	time.Sleep(1 * time.Second)
// 	// }
// 	var name string
// 	var age int

// 	fmt.Print("Enter your name and age: ")
// 	_, err := fmt.Scan(&name, &age)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

//		fmt.Printf("Name: %s, Age: %d\n", name, age)
//	}

package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// // 打开目录
	// dir, err := os.Open(".") // 打开当前目录
	// if err != nil {
	// 	fmt.Println("Error opening directory:", err)
	// 	return
	// }
	// defer dir.Close()

	// // 读取目录项
	// files, err := dir.Readdir(-1) // -1 表示读取所有目录项
	// if err != nil {
	// 	fmt.Println("Error reading directory:", err)
	// 	return
	// }

	// // 遍历目录项并打印
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		fmt.Printf("Directory: %s\n", file.Name())
	// 	} else {
	// 		fmt.Printf("File: %s\n", file.Name())
	// 	}
	// }
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to server: ", err)
	}
	defer conn.Close()

	// file, err := os.Open("/opt/sxs/shell/cargo-robot-1.1.5.deb")
	// if err != nil {
	// 	log.Fatal("Error opening file:", err)
	// }
	// defer file.Close()

	// fileinfo, err := file.Stat()
	// if err != nil {
	// 	log.Fatal("Error getting file info:", err)
	// }

	// buffer := make([]byte, fileinfo.Size())
	// _, err = file.Read(buffer)

	// if err != nil {
	// 	log.Fatal("Error reading file:", err)
	// }
	// conn.WriteMessage(websocket.BinaryMessage, buffer)

	// messageType, p, err := conn.ReadMessage()
	// fmt.Println(string(messageType))
	// fmt.Println(string(p))
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("readmessage is error ", err)
			break
		}
		fmt.Println(messageType)
		fmt.Println(string(msg))
	}
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// func main() {

// 	type Person struct {
// 		Name string `json:"name"`
// 		Age  int    `json:"age"`
// 	}

// 	// 创建一个 Person 实例
// 	p := Person{
// 		Name: "Alice",
// 		Age:  30,
// 	}

// 	err := json.NewEncoder(os.Stdout).Encode(p)
// 	if err != nil {
// 		fmt.Println("Error encoding JSON:", err)
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"selftest/model"
// 	"strconv"
// )

// func main() {
// 	model.Conn = model.NewConn()
// 	//设置默认路由和处理函数

// 	http.HandleFunc("/download", Login)
// 	fmt.Println("服务器运行的端口 8008")
// 	http.ListenAndServe(":8008", nil)
// }
// func Login(w http.ResponseWriter, r *http.Request) {
// 	//只允许post请求
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}
// 	//解析并处理post请求体中的数据
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "faild to parse form", http.StatusBadRequest)
// 		return
// 	}

// 	id := r.FormValue("id")
// 	name := r.FormValue("name")
// 	id64, _ := strconv.ParseInt(id, 10, 64)
// 	user := model.Conn.User(id64)

// 	if user.ID == id64 && user.Name != name {
// 		json.NewEncoder(w).Encode("sss")
// 		return
// 	} else if user.ID == id64 && user.Name == name {
// 		fmt.Fprintf(w, "登录成功")
// 	} else {
// 		fmt.Fprintf(w, "用户不存在")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// func Download(w http.ResponseWriter, r *http.Request) {
// 	//打开文件
// 	file, err := os.Open("/home/cargo/openssh-9.0p1.tar.gz")
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("unable to openfile: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	defer file.Close()
// 	//获取文件信息
// 	fileinfo, err := file.Stat()
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("unable to get file info: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	fileSize := fileinfo.Size()
// 	//检查是否range请求头
// 	rangeHeader := r.Header.Get("Range")
// 	if rangeHeader != "" {
// 		//解析range请求头
// 		rangeValue, err := parseRange(rangeHeader, fileSize)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("invalid range head:%v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		//设置响应头
// 		w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", rangeValue.start, rangeValue.end, fileSize))
// 		w.Header().Set("Content-Length", fmt.Sprintf("%d", rangeValue.length))
// 		w.WriteHeader(http.StatusPartialContent)
// 		//发送部分文件内容
// 		if _, err := file.Seek(rangeValue.start, 0); err != nil {
// 			http.Error(w, fmt.Sprintf("unable to seek file:%v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		if _, err := io.CopyN(w, file, int64(rangeValue.length)); err != nil {
// 			http.Error(w, fmt.Sprintf("unable to seek file:%v", err), http.StatusInternalServerError)
// 			return
// 		}
// 	} else {
// 		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))
// 		w.WriteHeader(http.StatusOK)
// 		if _, err := io.Copy(w, file); err != nil {
// 			http.Error(w, fmt.Sprintf("unable to copy file content %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }
// func parseRange(header string, filesize int64) (rangeValue, error) {
// 	const prefix = "bytes="
// 	if len(header) <= len(prefix) || header[:len(prefix)] != prefix {
// 		return rangeValue{}, fmt.Errorf("invalid range header")
// 	}
// 	var start, end int64
// 	var length int
// 	fmt.Sscanf(header[len(prefix):], "%d-%d", &start, &end)
// 	if end == 0 {
// 		end = filesize - 1
// 	}
// 	length = int(end - start + 1)
// 	return rangeValue{start: start, end: end, length: length}, nil
// }

// type rangeValue struct {
// 	start  int64
// 	end    int64
// 	length int
// }

// func main() {
// 	http.HandleFunc("/", Download)
// 	fmt.Println("server is running in http://localhost:8080")
// 	if err := http.ListenAndServe(":9999", nil); err != nil {
// 		fmt.Println("faild to start server", err)
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	router := mux.NewRouter()
// 	//开启协程启动connection服务管理中心
// 	go h.run()
// 	//创建ws服务
// 	router.HandleFunc("/ws", Myws)
// 	//启动http服务
// 	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
// 		fmt.Println("err:", err)
// 	}
// }

// // 用户中心，维护多个用户的connection
// var h = hub{
// 	c: make(map[*connection]bool),
// 	u: make(chan *connection),
// 	b: make(chan []byte),
// 	r: make(chan *connection),
// }

// type hub struct {
// 	//当前在线connection信息
// 	c map[*connection]bool
// 	//删除connection
// 	u chan *connection
// 	//传递数据
// 	b chan []byte
// 	//加入connection
// 	r chan *connection
// }

// func (h *hub) run() {
// 	for {
// 		select {
// 		//用户连接，添加connection信息
// 		case c := <-h.r:
// 			h.c[c] = true
// 			c.data.Ip = c.ws.RemoteAddr().String()
// 			c.data.Type = "handshake"
// 			c.data.UserList = user_list
// 			data_b, _ := json.Marshal(c.data)
// 			//发送给写入器
// 			c.sc <- data_b
// 		//删除指定用户连接
// 		case c := <-h.u:
// 			if _, ok := h.c[c]; ok {
// 				delete(h.c, c)
// 				close(c.sc)
// 			}
// 		//向聊天室在线人员发送信息
// 		case data := <-h.b:
// 			for c := range h.c {
// 				select {
// 				//发送数据
// 				case c.sc <- data:
// 				//发送不成功则删除connection信息
// 				default:
// 					delete(h.c, c)
// 					close(c.sc)
// 				}
// 			}
// 		}
// 	}

// }
