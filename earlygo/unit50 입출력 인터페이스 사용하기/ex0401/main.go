package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, world!"
	r := strings.NewReader(s)

	var s1, s2 string
	n, _ := fmt.Fscanf(r, "%s %s", &s1, &s2)

	fmt.Println("입력 개수 : ", n)
	fmt.Println(s1)
	fmt.Println(s2)
}
