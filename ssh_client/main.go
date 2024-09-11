package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// type PackageMessage struct {
// 	FileName string `json:"file_name"`
// 	Content  []byte `json:"content"`
// }

func main() {
	serverAddr := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Error connecting to server: ", err)
	}
	defer conn.Close()
	// Create file to save the received .deb file
	// var packageMessage PackageMessage
	// err = conn.ReadMessage(&packageMessage)
	// if err != nil {
	// 	log.Fatal("Error reading JSON message: ", err)
	// }
	// // Save received file
	// file, err := os.Create(packageMessage.FileName)
	// if err != nil {
	// 	log.Fatal("Error creating file: ", err)
	// }
	// defer file.Close()
	// _, err = file.Write(packageMessage.Content)
	// if err != nil {
	// 	log.Fatal("Error writing to file: ", err)
	// }

	// fmt.Printf("File %s received successfully.\n", packageMessage.FileName)
	// Read end of file signal
	// Receive .deb package from WebSocket server
	// _, msg, _ := conn.ReadMessage()
	// fmt.Println(string(msg))
	// _, file, err := conn.NextReader()
	// if err != nil {
	// 	log.Fatal("Error while receiving file:", err)
	// }
	for {
		data := []byte("hello client")
		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			fmt.Println("conn is error")
		}
		time.Sleep(time.Second)
	}
	// _, msg, err := conn.ReadMessage()
	// if err != nil {
	// 	log.Fatal("Error reading end of file message: ", err)
	// }
	// if string(msg) != "EOF" {
	// 	log.Fatal("Unexpected end of file message.")
	// }

	// // Install the .deb package
	// cmd := exec.Command("dpkg", "-i", packageMessage.FileName)
	// var out bytes.Buffer
	// cmd.Stdout = &out
	// cmd.Stderr = &out
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatalf("Error installing .deb package: %v\nOutput: %s", err, out.String())
	// }
	// fmt.Println("Package installed successfully.")
	// Save the file
	// outFile, err := os.Create("/tmp/package.deb")
	// if err != nil {
	// 	log.Fatal("Error while creating file:", err)
	// }
	// defer outFile.Close()
	// _, err = io.Copy(outFile, file)
	// if err != nil {
	// 	log.Fatal("Error while saving file:", err)
	// }

	// // Install the .deb package
	// cmd := exec.Command("dpkg", "-i", "/tmp/package.deb")
	// err = cmd.Run()
	// if err != nil {
	// 	log.Fatal("Error while installing package:", err)
	// }
	// fmt.Println("Package installed successfully.")
}
