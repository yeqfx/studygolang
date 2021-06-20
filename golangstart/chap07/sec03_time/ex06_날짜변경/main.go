package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	fmt.Println(t0)

	t1 := t0.AddDate(1, 0, 0)
	fmt.Println(t1)

	t2 := t0.AddDate(0, -1, 0)
	fmt.Println(t2)

	t3 := t0.AddDate(0, 0, -1)
	fmt.Println(t3)

	t4 := t0.AddDate(0, 2, -12)
	fmt.Println(t4)
}
