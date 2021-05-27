package main

import "fmt"

func main() {
	c := make(chan int, 1)

	go func() {
		c <- 1
	}()

	a, ok := <-c
	fmt.Println(a, ok)

	close(c)

	a, ok = <-c
	fmt.Println(a, ok)
}
