package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	fmt.Println(os.Getuid())
	fmt.Println(os.Geteuid())
	fmt.Println(os.Getgid())
	fmt.Println(os.Getegid())
}
