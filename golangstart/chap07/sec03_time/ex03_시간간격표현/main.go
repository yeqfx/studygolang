package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)

	t := time.Now()
	fmt.Println(t)
	t = t.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t)
}
