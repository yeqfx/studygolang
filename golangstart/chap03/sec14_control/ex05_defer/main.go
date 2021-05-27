package main

import (
	"fmt"
	"log"
	"os"
)

func runDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("done")
}
func main() {
	runDefer()

	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatal("Error!")
	}
	fmt.Println("File Opened!!")
	defer file.Close()

	defer func() {
		fmt.Println("A")
		fmt.Println("B")
		fmt.Println("C")
	}()
}
