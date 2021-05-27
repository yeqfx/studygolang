package main

import "fmt"

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func main() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	fmt.Printf("%t, %T, %p, %P\n", err, err, err, err)
	panic(err)
}
