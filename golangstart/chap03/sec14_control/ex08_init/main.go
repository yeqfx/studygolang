package main

import "fmt"

func main() {
	fmt.Println(S)
}

var S = ""

func init() {
	S = S + "A"
}

func init() {
	S = S + "B"
}

func init() {
	S = S + "C"
}
