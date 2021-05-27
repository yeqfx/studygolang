package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	if i, ok := <-ch; ok {
		fmt.Println(ok)
		fmt.Println(i)
	} else {
		fmt.Println(ok)
	}
	if i, ok := <-ch; ok {
		fmt.Println(ok)
		fmt.Println(i)
	} else {
		fmt.Println(ok)
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
