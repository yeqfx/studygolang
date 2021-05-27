package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

type IntPoint struct{ X, Y int }
type FloatPoint struct{ X, Y float64 }

func (p *IntPoint) Distance(dp *IntPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	p := &Point{5, 12}
	p.Render()
	fmt.Println("")

	p1 := &Point{X: 0, Y: 0}
	distance := p1.Distance(&Point{X: 1, Y: 1})
	fmt.Println(distance)
	fmt.Println("")

	pf := new(FloatPoint)
	pf.X = 1.0
	pf.Y = 1.0
	distance = pf.Distance(&FloatPoint{X: 2.0, Y: 2.0})
	fmt.Println(distance)
}
