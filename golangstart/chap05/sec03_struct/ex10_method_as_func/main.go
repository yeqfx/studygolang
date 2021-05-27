package main

import "fmt"

type Point struct{ X, Y int }

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d, %d]", p.X, p.Y)
}

func main() {
	f := (*Point).ToString
	fp := f(&Point{X: 7, Y: 11})
	fmt.Println(fp)

	fp = ((*Point).ToString(&Point{X: 11, Y: 33}))
	fmt.Println(fp)

	p := &Point{X: 2, Y: 3}
	fs := p.ToString
	fp = fs()
	fmt.Println(fp)
}
