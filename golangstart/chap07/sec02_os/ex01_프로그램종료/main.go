package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// defer는 os.Exit에 의해 무시된다.
	defer func() {
		fmt.Println("defer")
	}()
	_, err := os.Open("foo")
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
