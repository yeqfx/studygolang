package hangul

import (
	"fmt"

	"../../pkg/hangul"
)

func ExampleHasConsonantSuffix() {
	fmt.Println(hangul.HasConsonantSuffix("Go 언어"))
	fmt.Println(hangul.HasConsonantSuffix("그럼"))
	fmt.Println(hangul.HasConsonantSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	// true
	// false
}

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	//Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printByte2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	//Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifyBytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}
