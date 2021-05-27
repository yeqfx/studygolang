package main

import (
	"fmt"

	"./foo"
)

func main() {
	fmt.Println(foo.MAX)
	// foo.internal_const

	fmt.Println(foo.FooFunc(5))
	// foo.internalFunc(5)
}
