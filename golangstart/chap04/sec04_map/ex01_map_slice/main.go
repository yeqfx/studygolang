package main

import "fmt"

func main() {
	m := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
		4: nil,
	}

	s1 := m[4]
	if s1 != nil {
		fmt.Println(s1)
		fmt.Println("s1 is not nil")
	}

	if s2, ok := m[4]; ok {
		fmt.Println(s2)
		fmt.Println("s2 is", ok)
	}
}
