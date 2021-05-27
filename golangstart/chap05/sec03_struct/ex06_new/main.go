package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Area string
}

type Point struct {
	X, Y int
}

func main() {

	p := new(Person)
	fmt.Printf("type : %T, address : %p\n", p, p)

	i := new(int)
	fmt.Printf("type : %T, address : %p\n", i, i)

	sa := new([]string)
	fmt.Printf("type : %T, address : %p\n", sa, sa)

	s := new(string)
	fmt.Printf("type : %T, address : %p\n", s, s)

	fmt.Println("")
	// new를 사용하여 구조체 포인터 초기화
	pp := new(Point)
	pp.X = 1
	pp.Y = 2
	fmt.Println(pp, *pp)
	// 복합 문법을 사용하여 구조체 포인터 초기화
	ppp := &Point{X: 1, Y: 1}
	fmt.Println(ppp, *ppp)

	pppp := new(Point)
	fmt.Println(pppp, *pppp)
}
