package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
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

	s := "Hello, world!"
	r := strings.NewReader(s)

	w := bufio.NewWriter(file)
	w.ReadFrom(r)
	w.Flush()

	io.Copy(w, r)
	w.Flush()
}
