package write

import (
	"fmt"
)

// var lock sync.Mutex

func write() {
	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	// for {
	// 	<-exitChan
	// 	fmt.Println("结束")
	// 	break
	// }
	data, ok := <-exitChan
	if !ok {
		return
	}
	fmt.Println(data)
}

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("写入数据", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取到:", val)
		//time.Sleep(time.Second)

	}
	exitChan <- true
	//close(exitChan)
}
