package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 3)
	ch1 <- 1
	ch2 <- 2

	select {
	case <-ch1:
		fmt.Println("ch1에서 수신")
	case <-ch2:
		fmt.Println("ch2에서 수신")
	case ch3 <- 3:
		fmt.Println("ch3로 발신")
	default:
		fmt.Println("처리되지 않음")
	}
}
