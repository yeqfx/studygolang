package main

import "fmt"

func sum(a, b int, c chan int) {
	c <- a + b
}

func main() {
	c := make(chan int)

	go sum(1, 2, c)

	n := <-c
	fmt.Println(n)
}
