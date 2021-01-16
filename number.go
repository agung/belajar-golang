package main

import "fmt"

// deklarasi type data int8 (8 bytes)
var (
	number8 int8 	// max and min = 127, -128
	number16 int16	// max and min = 32767, -32768
	number32 int32	// max and min = 2147483647, -2147483648
	number64 int64	// max and min = 9223372036854775807, -9223372036854775808

	// float check in docs
)


func main() {
	number8 = 100
	fmt.Println(number8)
}
