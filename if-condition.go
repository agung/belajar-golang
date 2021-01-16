package main

import (
	"fmt"
)

func main() {
	// short condition
	name := "Agung Nugraha"
	if len(name) > 5 {
		fmt.Println("Nama terlalu panjang")
	}

	if length := len(name); length > 5 {
		fmt.Printf("Nama %s panjangnya %d", name, length)
	}
}
