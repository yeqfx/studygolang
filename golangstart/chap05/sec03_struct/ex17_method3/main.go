package main

import "fmt"

type MySlice []interface{}

// func NewMySlice

func (m MySlice) MyAppend(i interface{}) MySlice {
	return append(m, i)
}

func (m MySlice) MyAppend2(i ...interface{}) MySlice {
	return append(m, i)
}

func main() {
	var mySlice = MySlice{}
	// mySlice = make(MySlice, 0)
	mySlice = mySlice.MyAppend(1)
	fmt.Println(mySlice)
}
