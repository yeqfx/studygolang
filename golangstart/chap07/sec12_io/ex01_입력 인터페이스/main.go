package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
