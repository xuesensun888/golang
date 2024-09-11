package handler

import (
	"camera/src/shell"
	"fmt"
)

func Getopic() {
	fmt.Println(111)
	stdout, _, _, pid := shell.Command("ls")
	fmt.Println(stdout)
	fmt.Println(pid)
}
