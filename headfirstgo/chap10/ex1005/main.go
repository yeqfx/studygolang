package main

import (
	"fmt"
	"log"

	"../ex1003/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(1967)
	if err != nil {
		log.Fatal(err)
	}
	if err := event.SetMonth(3); err != nil {
		log.Fatal(err)
	}
	if err := event.SetDay(1); err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}
