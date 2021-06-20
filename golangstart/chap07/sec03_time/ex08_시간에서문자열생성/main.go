package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.RFC3339Nano))
	fmt.Println(t.Format("2006년1월2일 15시04분05초"))
}
