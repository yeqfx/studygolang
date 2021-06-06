package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args의 크기를 표시
	fmt.Printf("length=%d\n", len(os.Args))
	// os.Args의 내용을 모두 출력
	for _, v := range os.Args {
		fmt.Println(v)
	}
}
