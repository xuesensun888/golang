package test

import (
	"fmt"
	"time"
	//"time"
)

func test() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)
	start := time.Now().Unix()
	//开启一个携程 向intchan放入数据
	go putNum(intChan)
	//go readNum(intChan,exitChan)
	//开启四个携程 从intchan取出数据。并判断是否为素数，如果是 就放入prime
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	//z这里开启主线程 进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用携程耗时==", end-start)
		close(primeChan)
	}()
	time.Sleep(time.Second)
	//<- exitChan
	//for {
	//	_, ok := <- primeChan
	//	if !ok{
	//		break
	//	}

	//fmt.Printf("素数=%d\n", v)
	//}
	fmt.Println("main主线程退出")
}

// 向intchan放入8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 800; i++ {
		intChan <- i
	}
	fmt.Println("intchan xirru")
	close(intChan)
}
func readNum(intchan chan int, exitChan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			break
		}
		fmt.Printf("readdata duqudapdeshuju=%v\n", v)
	}
	exitChan <- true
	close(exitChan)
}

// c从intchan取出数据并判断是否为素数，如果是就放入到primechan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		//判断是否为素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("有一个primenum协成因为取不到数据退出了")
	exitChan <- true
}
