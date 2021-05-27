package main

import "fmt"

type IntPair []int

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Second() int {
	return ip[1]
}

type Strings []string

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return sum
}

func main() {
	ip := IntPair{1, 2}
	fmt.Println(ip.First())
	fmt.Println(ip.Second())

	fmt.Println(Strings{"A", "B", "C"}.Join(","))
	fmt.Println(Strings{"E", "F", "G"})

	// var s []string {"I", "J", "k"}
	s := []string{"I", "J", "K"}
	fmt.Println(s)
}
