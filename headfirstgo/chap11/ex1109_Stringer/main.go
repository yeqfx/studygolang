package main

import "fmt"

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func main() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
	fmt.Print(coffeePot, "\n")
	fmt.Println(coffeePot)
	fmt.Printf("%s", coffeePot)
}
