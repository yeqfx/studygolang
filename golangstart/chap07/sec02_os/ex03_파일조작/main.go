package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, _ := os.Create("foo.txt")

	fi, _ := f.Stat()
	fmt.Println(fi.Name())
	fmt.Println(fi.Size())
	fmt.Println(fi.IsDir())

	f.Write([]byte("Hello, World!\n"))
	f.WriteAt([]byte("Golang"), 7)
	f.Seek(0, io.SeekEnd)
	f.WriteString("Yeah!")

	f.Seek(0, io.SeekStart)
	bs := make([]byte, 128)
	f.Read(bs)

	fmt.Println(string(bs))

	_, err := f.Seek(-8, io.SeekCurrent)
	if err != nil {
		log.Fatal(err)
	}
	bs = make([]byte, 128)
	f.Read(bs)
	fmt.Println(string(bs))
}
