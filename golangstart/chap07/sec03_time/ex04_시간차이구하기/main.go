package main

import (
	"fmt"
	"time"
)

func main() {
	// 내년 내 생일까지 남은 시간
	t0 := time.Date(2022, 8, 21, 0, 0, 0, 0, time.Local)
	t1 := time.Now()
	d := t0.Sub(t1)
	fmt.Println(d)
}
