package main

import (
	"fmt"

	"../magazine"
)

func main() {
	subscriber := magazine.Subscriber{Rate: 4.99}
	fmt.Println(subscriber)
}
