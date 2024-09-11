package main

import (
	"fmt"
	"kargo-deploy/service"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// wg.Add(1)
	// go service.ListenAndServSocket(&wg)
	// //time.Sleep(100 * time.Second)
	// wg.Wait()
	fmt.Println("开始")
	wg.Add(1)
	go service.ListenAndServSocket(&wg)
	time.Sleep(time.Second * 1)
	wg.Add(1)
	go service.Update(&wg)
	wg.Wait()

}
