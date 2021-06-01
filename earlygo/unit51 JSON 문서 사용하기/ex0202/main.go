package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommands []string  `json:"recommands"`
	Comments   []Comment `json:"comments"`
}

func main() {
	b, err := ioutil.ReadFile("./article.json")
	if err != nil {
		log.Fatal(err)
	}

	var data []Article
	json.Unmarshal(b, &data)

	fmt.Println(data)
}
