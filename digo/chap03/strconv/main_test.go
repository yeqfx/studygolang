package main

import (
	"fmt"
	"strconv"
)

func Example_main() {
	var i int
	var k int64
	var f float64
	var s string
	var err error

	if i, err = strconv.Atoi("350"); err == nil {
		fmt.Printf("%T, %v\n", i, i)
	}

	if k, err = strconv.ParseInt("cc7fdd", 16, 32); err == nil {
		fmt.Printf("%T, %v\n", k, k)
	}

	if k, err = strconv.ParseInt("0xcc7fdd", 0, 32); err == nil {
		fmt.Printf("%T, %v\n", k, k)
	}

	if f, err = strconv.ParseFloat("3.14", 64); err == nil {
		fmt.Printf("%T, %f\n", f, f)
	}

	s = strconv.Itoa(340)
	fmt.Printf("%T, %v\n", s, s)

	s = strconv.FormatInt(13402077, 16)
	fmt.Printf("%T, %v\n", s, s)

	// Output:
	// int, 350
	// int64, 13402077
	// int64, 13402077
	// float64, 3.140000
	// string, 340
	// string, cc7fdd

}
