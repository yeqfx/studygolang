package main

import "fmt"

type hello interface{}

func main() {
	var h hello
	fmt.Println(h)
}
