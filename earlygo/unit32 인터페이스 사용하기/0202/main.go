package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch v := arg.(type) {
	case Person:
		return v.name + " " + strconv.Itoa(v.age)
	case *Person:
		return v.name + " " + strconv.Itoa(v.age)
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(Person{"Maria", 29}))
	fmt.Println(formatString(&Person{"John", 12}))

	var andrew = new(Person)
	andrew.name = "Andrew"
	andrew.age = 35

	fmt.Println(formatString(andrew))

	var t interface{} = Person{"Maria", 20}

	if v, ok := t.(Person); ok {
		fmt.Println(v, ok)
	}
}
