package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t, err := time.Parse("2006/01/02", "2015/11/27")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	t, err = time.Parse(time.RFC822, "27 Nov 15 18:00 KST")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	t, err = time.Parse("2006년 1월 2일 15시 04분 05초", "2015년 11월 27일 14시 30분 29초")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	t, err = time.Parse("2006", "2015")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Zone())
}
