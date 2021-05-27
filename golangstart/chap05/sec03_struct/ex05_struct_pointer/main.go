package main

import "fmt"

type Point struct {
	X, Y int
}

func swap(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func main() {
	p := Point{X: 1, Y: 2}
	swap(&p)

	fmt.Println(p.X)
	fmt.Println(p.Y)
}
