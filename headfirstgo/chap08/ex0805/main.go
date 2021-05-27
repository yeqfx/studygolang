package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.name)

	fmt.Println("")
	subscriber3 := defaultSubscriber("Aman singh")
	subscriber3.rate = 4.99
	printInfo(subscriber3)
	subscriber4 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber4)

	fmt.Println("")
	var subscriber5 subscriber
	applyDiscount(&subscriber5)
	fmt.Println(subscriber5.rate)

}
