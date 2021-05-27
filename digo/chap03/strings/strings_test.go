package main

import (
	"fmt"
	"strings"
)

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)

	// Output:
	// abcdef
	// abcdef
}

func Example_sPrint() {
	s := "abc"
	s1 := fmt.Sprint(s, "def")
	fmt.Println(s1)
	s2 := fmt.Sprintf("%sdef", s)
	fmt.Println(s2)
	s3 := strings.Join([]string{s, "def"}, "")
	fmt.Println(s3)
	ps := &s3
	fmt.Println(*ps)
	//output:
	//abcdef
	//abcdef
	//abcdef
	//abcdef
}
