package main

import (
	"fmt"
	"os"
)

func main() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
	fmt.Println("------------------------")
	fmt.Println(os.Getenv("HOME"))
	os.Setenv("A", "B")
	fmt.Println(os.Getenv("A"))
	os.Setenv("A", "C")
	fmt.Println(os.Getenv("A"))
	os.Setenv("A", "1")
	fmt.Println(os.Getenv("A"))
	os.Unsetenv("A")
	fmt.Println(os.Getenv("A"))
	if home, ok := os.LookupEnv("HOME"); ok {
		fmt.Println(home)
	} else {
		fmt.Println("no $HOME")
	}
}
