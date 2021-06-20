package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2015, 11, 27, 15, 0, 0, 0, time.UTC)
	fmt.Println(t)
	KST := t.Local()
	fmt.Println(KST)
}
