package main

import "fmt"

func main() {
	for i, r := range "ABC가나다" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}
}
