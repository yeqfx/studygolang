package main

import "fmt"

type Point struct {
	X, Y int
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}
func main() {
	s := struct{ X, Y int }{X: 1, Y: 2}
	showStruct(s)
	p := Point{X: 3, Y: 8}
	showStruct(p)
}
