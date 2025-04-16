package main

import (
	"fmt"
	"os"
)

func main() {
	name, _ := os.Hostname()
	// 打印主机名
	fmt.Printf("%s\n", name)
}
