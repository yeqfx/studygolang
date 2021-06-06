package main

import (
	"fmt"
	"os"
)

func CloseFile(f *os.File) {
	// fmt.Println("파일을 닫습니다.")
	f.Close()
}

func main() {
	if dir, err := os.Getwd(); err == nil {
		fmt.Println(dir)
	}

	if err := os.Chdir(".."); err == nil {
		if dir, err := os.Getwd(); err == nil {
			fmt.Println(dir)
		}
	}
	if f, err := os.Open("."); err == nil {
		defer CloseFile(f)
		if fis, err := f.Readdir(0); err == nil {
			for _, fi := range fis {
				if fi.IsDir() {
					fmt.Println(fi.Name())
				}
			}
		}
	}
	if err := os.Chdir("./ex04_디렉토리조작"); err == nil {
		if dir, err := os.Getwd(); err == nil {
			fmt.Println(dir)
		}
	}
	// 디렉토리 만들기
	// err := os.Mkdir("foo", 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := os.MkdirAll("foo/bar/baz", 0775); err == nil {
		fmt.Println("디렉토리 foo/bar/baz 생성")
	}

	if err := os.RemoveAll("foo"); err == nil {
		fmt.Println("디렉토리 foo/bar/baz 삭제")
	}
	// 시스템 임시 디렉토리 얻기
	fmt.Println(os.TempDir())
	// 심볼릭 링크 얻기
	if err := os.Symlink("main.go", "foo.go"); err == nil {
		if path, err := os.Readlink("foo.go"); err == nil {
			fmt.Println(path)
			os.Remove("foo.go")
		}
	}
	// // 호스트 이름 검색
	if host, err := os.Hostname(); err == nil {
		fmt.Println(host)
	}
}
