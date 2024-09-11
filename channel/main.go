package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Month()
	fmt.Println(weekday)
	fmt.Println(time.Thursday)
}
