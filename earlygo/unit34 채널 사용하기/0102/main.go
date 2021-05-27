package main

import (
	"fmt"
	"time"
)

func goroutine(count int, done chan bool) {
	for i := 0; i < count; i++ {
		done <- true
		fmt.Println("고루틴 : ", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	done := make(chan bool)
	count := 3

	go goroutine(count, done)

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수 : ", i)
	}
}
