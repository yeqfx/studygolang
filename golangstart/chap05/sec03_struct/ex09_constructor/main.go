package main

import "fmt"

type User struct {
	id   int
	name string
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.id = id
	u.name = name
	return u
}

func main() {
	fmt.Println(NewUser(1, "Taro"))
	u := NewUser(2, "Taro2")
	fmt.Println(u.id)
	fmt.Println(u.name)
}
