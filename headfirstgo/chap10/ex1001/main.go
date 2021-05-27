package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}
	// err := date.SetYear(2021)
	if err := date.SetYear(2021); err != nil {
		log.Fatal(err)
	}
	if err := date.SetMonth(5); err != nil {
		log.Fatal(err)
	}
	if err := date.SetDay(15); err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
}
