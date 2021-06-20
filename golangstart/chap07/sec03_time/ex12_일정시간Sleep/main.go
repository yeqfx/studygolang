package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	fmt.Println("5초간 정지 후 프린트되었음")
}
