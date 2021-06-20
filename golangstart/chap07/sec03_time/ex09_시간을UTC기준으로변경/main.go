package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	utc := t.UTC()
	fmt.Println(utc)
	local := t.Local()
	fmt.Println(local)
}
