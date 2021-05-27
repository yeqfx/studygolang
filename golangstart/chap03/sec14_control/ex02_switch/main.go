package main

import "fmt"

func main() {
	switch b := true; b {
	case true, false:
		fmt.Println("true or false")
		// case b != false:
		// fmt.Println("is not false")
	}
}
