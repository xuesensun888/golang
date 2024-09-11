package internal

import (
	"fmt"
	"sync"
)

func Waiting(wg *sync.WaitGroup) {
	fmt.Println("entry waiting....")
	wg.Wait()
}
