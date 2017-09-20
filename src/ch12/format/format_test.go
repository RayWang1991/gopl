package format

import "fmt"

func ExampleAny() {
	values := []interface{}{
		// numbers
		1,
		2,
		1.23,
		-100e3,
		77 + 36.3i,
		// strings
		"string",
		"中国",
		// bool
		true,
		false,
		// struct
		struct{}{},
		struct{ int }{5},
	}
	for _, v := range values {
		fmt.Println(Any(v))
	}
	// Output:
	// 1
	// 2
	// 1.23
	// -100000
	// 77+36.3i
	// "string"
	// "中国"
	// true
	// false
	// struct {}value
	// struct { int }value
}
