package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	unix := t.Unix()
	fmt.Println(unix)

	unix_t := time.Unix(1624201042, 0)
	fmt.Println(unix_t)
}
