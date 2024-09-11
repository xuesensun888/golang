package main

import (
	"flag"
	"fmt"
	"time"
)

var name = flag.String("qwe", "sxs", "go flag test")

func main() {
	flag.Set("qwe", "kkkk")
	result := flag.Parsed()
	fmt.Println(result)
	flag.Parse()
	result1 := flag.Parsed()
	fmt.Println(result1)
	fmt.Println("hello", *name)
	now := time.Now().Weekday()
	fmt.Println(now)
	// var weekDay int = now
	a := time.Saturday
	fmt.Printf("%T\n", a)
	// fmt.Println(weekDay)
	fmt.Println(int(a))
	fmt.Println(2 * time.Second)
	fmt.Println(int(time.Saturday))
	fmt.Println(time.Month(time.February))
	fmt.Println(time.Now().Date())

}
