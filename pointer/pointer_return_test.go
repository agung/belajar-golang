package pointer

import (
	"fmt"
	"testing"
)

func number() *int {
	num := 100
	return &num
}

func TestNumber(t *testing.T) {
	// jika tidak menggunakan * makan return pointer address
	fmt.Println(*number())
}
