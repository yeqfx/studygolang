package main

import "fmt"

type Point struct{ X, Y int }

type Points []*Point

func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}
	return str
}

func main() {
	p := make([]Point, 2)
	fmt.Printf("%T, %v, %#v\n", p, p, p)
	for _, p := range p {
		fmt.Println(p.X, p.Y)
	}

	ps1 := make(Points, 0)
	fmt.Printf("%T, %v, %#v\n", ps1, ps1, ps1)

	ps := Points{}
	fmt.Printf("%T, %v, %#v\n", ps, ps, ps)
	ps = append(ps, &Point{X: 1, Y: 2})
	ps = append(ps, nil)
	ps = append(ps, &Point{X: 3, Y: 4})
	fmt.Println(ps.ToString())

}
