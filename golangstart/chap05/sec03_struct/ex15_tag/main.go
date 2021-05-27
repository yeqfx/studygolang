package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `UserID`
	Name string `name`
	Age  uint   `age`
}

func main() {
	u := User{Id: 1, Name: "Taro", Age: 32}
	t := reflect.TypeOf(u)
	fmt.Printf("%T\n", t)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}
}
