package main

import "fmt"

type MyType string

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) WithReturn() int {
	return len(m)
}

func main() {
	value := MyType("MyType value")
	value.MethodWithParameters(4, true)

	lenvalue := value.WithReturn()
	fmt.Println(lenvalue)
}
