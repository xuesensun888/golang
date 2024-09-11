package main

import (
	"fmt"
	"time"
)

func context() {
	//创建一个父context
	parentCtx := context.Background()
	//创建一个带有取消函数的子级 context 设置超时为两秒
	childCtx, _ := context.WithTimeout(parentCtx, 2*time.Second)
	//defer cancel()

	//在子级的context中启动一个gorouting
	go doSomething(childCtx)

	//阻塞一段时间
	time.Sleep(10 * time.Second)
	//调用取消函数 取消子级context中的操作
	fmt.Println("调用撤销函数")

	//阻塞一段时间
	time.Sleep(1 * time.Second)
}

func doSomething(ctx context.Context) {
	//模拟一个耗时操作
	for {
		select {
		case <-ctx.Done():
			//当context被取消时停止操作
			fmt.Println("operation canceled.")
			return
		default:
			fmt.Println("performing operation...")
			time.Sleep(500 * time.Microsecond)
		}
	}
}
