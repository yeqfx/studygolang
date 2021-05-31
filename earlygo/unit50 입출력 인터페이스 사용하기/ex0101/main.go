package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Printf("%T, %T\n", file, w)

	w.WriteString("Hello, world!")
	w.Flush()

	r := bufio.NewReader(file)

	fi, _ := file.Stat()
	b := make([]byte, fi.Size())

	file.Seek(0, io.SeekStart)
	r.Read(b)

	fmt.Println(string(b))
}
