package main

import "fmt"

func printString(s string) {
	fmt.Println(s)
	fmt.Printf("%T, %p", s, &s)
}

func main() {
	s := "ABC"
	fmt.Println(&s)
	fmt.Println(s[0])
	// fmt.Println(&s[0])

	printString(s)
}
