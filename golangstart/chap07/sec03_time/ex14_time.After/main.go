package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.After(5 * time.Second)
	fmt.Println(<-ch)
}
