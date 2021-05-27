package main

import (
	"fmt"
	"log"

	"./calendar"
)

func main() {
	date := calendar.Date{}
	if err := date.SetYear(2021); err != nil {
		log.Fatal(err)
	}
	if err := date.SetMonth(5); err != nil {
		log.Fatal(err)
	}
	if err := date.SetDay(19); err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	fmt.Println("")

	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
