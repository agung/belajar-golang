package main

import "fmt"

// membuat variable di luar dari method
// agar dapat dipanggil di semua method
var data string

func main() {
	data = "wkwkw"
	fmt.Println(data)

	data = "Agung Nugraha"
	fmt.Println(data)
}
