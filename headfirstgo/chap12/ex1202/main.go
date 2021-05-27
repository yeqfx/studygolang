package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("i don't want to talk")
	fmt.Println("Nice weather, eh?")
	return nil
}

func main() {

	if err := Socialize(); err != nil {
		log.Fatal()
	}
}
