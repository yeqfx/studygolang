package main

import "fmt"

type Feed struct {
	Name   string
	Amount int
}

type Animal struct {
	Name string
	Feed
}

func main() {
	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	fmt.Println(a.Name)
	fmt.Println(a.Feed.Name)
	fmt.Println(a.Amount)
}
