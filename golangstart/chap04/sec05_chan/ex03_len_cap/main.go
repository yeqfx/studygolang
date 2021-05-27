package main

import "fmt"

func main() {
	ch := make(chan string, 10)

	ch <- "Apple"
	fmt.Println(len(ch), cap(ch))
	ch <- "Banana"
	ch <- "Cherry"
	fmt.Println(len(ch), cap(ch))
}
