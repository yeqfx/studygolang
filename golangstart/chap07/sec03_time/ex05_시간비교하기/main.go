package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.Local)
	t1 := time.Date(2015, 11, 1, 0, 0, 0, 0, time.Local)

	fmt.Println(t1.Before(t0))
	fmt.Println(t1.After(t0))

	t2 := time.Date(2015, 10, 1, 9, 0, 0, 0, time.Local)
	t3 := time.Date(2015, 10, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t2.Equal(t3))
}
