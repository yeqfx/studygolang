package main

import "fmt"

type Point struct{ X, Y int }

// func (p Point) Set(x, y int) {
// 	p.X = x
// 	p.Y = y
// }

func (p *Point) Set(x, y int) {
	p.X = x
	p.Y = y
}

func main() {
	p1 := Point{}
	p1.Set(1, 2)
	fmt.Println(p1.X, p1.Y)

	p2 := &Point{}
	p2.Set(1, 2)
	fmt.Println(p2.X, p2.Y)
}
