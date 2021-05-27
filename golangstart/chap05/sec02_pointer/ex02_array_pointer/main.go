package main

import "fmt"

func main() {
	p := &[3]int{1, 2, 3}
	fmt.Println((*p)[0])
	fmt.Println(p[1])
	fmt.Println(p[2])
	fmt.Println("")

	fmt.Println(len(p))
	fmt.Println(cap(p))
	fmt.Println(p[0:2])

	sp := &[3]string{"Apple", "Banana", "Cherry"}

	for i, v := range sp {
		fmt.Println(i, v)
	}

	fmt.Printf("type = %T, address = %p\n", p, p)
	fmt.Printf("type = %T, address = %p\n", sp, sp)
}
