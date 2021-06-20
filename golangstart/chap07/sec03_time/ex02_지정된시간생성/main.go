package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t1 := time.Date(2021, 06, 20, 16, 43, 0, 0, time.Local)
	t2 := time.Date(2021, 06, 20, 16, 43, 0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t.Local())
	fmt.Println(t.Zone())
	fmt.Println(t.YearDay())
	fmt.Println(t.Month())
	fmt.Println(t.Weekday())
	fmt.Println(t.Nanosecond())
}
