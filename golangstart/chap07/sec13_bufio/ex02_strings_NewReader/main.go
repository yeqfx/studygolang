package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := `XXXXX
YYYYY
ZZZZZ`

	r := strings.NewReader(s)
	fmt.Printf("%T\n", r)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
	scanner.Scan()
	fmt.Println(scanner.Text())
}
