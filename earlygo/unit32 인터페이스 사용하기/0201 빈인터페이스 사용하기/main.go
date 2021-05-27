package main

import (
	"fmt"
	"strconv"
)

func formatString(arg interface{}) string {
	// switch arg.(type)
	switch v := arg.(type) {
	case int:
		// i := arg.(int)
		return strconv.Itoa(v)
	case float32:
		// f := arg.(float32)
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		// f := arg.(float64)
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		// s := arg.(string)
		return v
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println("Hello, world!")
}
