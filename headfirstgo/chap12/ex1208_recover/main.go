package main

import "fmt"

func calmDown() {
	fmt.Println(recover())
}

func freakOut() {
	defer calmDown()
	panic("oh no")
}

func main() {

	fmt.Println(recover())
	freakOut()
	fmt.Println("Exiting normally")
}
