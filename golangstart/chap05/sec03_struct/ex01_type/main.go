package main

import "fmt"

func main() {
	type MyInt int

	var n1 MyInt = 5
	n2 := MyInt(7)
	n3 := int(9)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
