package main

import "fmt"

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100

	// fmt.Println(<-c)
}

func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)

		fmt.Println(<-c)
	}
	// c <- 1
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
