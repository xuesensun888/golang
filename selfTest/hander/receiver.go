package hander

import "fmt"

func Receiver(c chan string) {
	for {
		msg := <-c // 从 channel 接收数据
		fmt.Println(msg)
	}
}
