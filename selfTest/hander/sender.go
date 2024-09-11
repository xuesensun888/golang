package hander

import (
	"fmt"
	"time"
)

func Sender(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("hello %d", i) // 向 channel 发送数据
		time.Sleep(time.Second)
	}
}
