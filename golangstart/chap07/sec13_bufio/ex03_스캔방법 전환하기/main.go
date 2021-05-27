package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := `ABC DEF
GHI JKL MNO
PQR STU VWX
YZ
한글 사랑
`
	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
