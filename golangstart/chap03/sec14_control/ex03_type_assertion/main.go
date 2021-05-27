package main

import "fmt"

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything(1)
	anything(3.14)
	anything(4 + 5i)
	anything('A')
	anything("한국어")
	anything([...]int{1, 2, 3, 4, 5})

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i)

	var y interface{} = 3.14
	j, isInt := y.(int)
	fmt.Println(j, isInt)
	f, isFloat64 := y.(float64)
	fmt.Println(f, isFloat64)
	s, isString := y.(string)
	fmt.Println(s, isString)

	var z interface{}
	if z == nil {
		fmt.Println("z is nil")
	} else if l, isInt := x.(int); isInt {
		fmt.Printf("z is integer : %d\n", l)
	} else if m, isString := x.(string); isString {
		fmt.Println(m)
	} else {
		fmt.Println("unsupported type!")
	}
}
