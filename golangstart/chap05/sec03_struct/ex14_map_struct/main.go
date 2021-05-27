package main

import "fmt"

type User struct {
	id   int
	name string
}

func main() {
	m1 := map[User]string{
		{id: 1, name: "Taro"}: "Tokyo",
		{id: 2, name: "Jiro"}: "Osaka",
	}
	m2 := map[int]User{
		1: {id: 1, name: "Taro"},
		2: {id: 2, name: "Jiro"},
	}
	fmt.Println(m1[User{1, "Taro"}])
	fmt.Println(m2)
	fmt.Println(m2[1])

	// 값이 슬라이스의 맵
	ms := map[int][]string{
		1: {"A", "B", "C"},
	}
	fmt.Println(ms)
	// 값이 맵의 맵
	mm := map[int]map[int]string{
		1: {1: "Apple", 2: "Banana", 3: "Cheery"},
	}
	fmt.Println(mm)
}
