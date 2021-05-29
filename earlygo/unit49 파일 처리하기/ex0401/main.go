package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	s := "Hello, world!"

	err := ioutil.WriteFile("hello.txt", []byte(s), os.FileMode(0644))

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
