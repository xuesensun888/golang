package test

import (
	"fmt"
	"time"
)

func test() {
	//构建一个通道
	ch := make(chan int)
	//开启一个并发的匿名函数
	go func() {
		fmt.Println("start gorountine")
		//通过通道通知main的grountine
		time.Sleep(time.Second * 5)
		ch <- 10
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait gorountine")
	data := <-ch

	fmt.Println(data)
	fmt.Println("all done")
}
