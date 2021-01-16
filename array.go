package main

import "fmt"

func main() {
	var nama [2]string
	nama[0] = "Agung"
	nama[1] = "Nugraha"

	fmt.Println(nama[0] + " " + nama[1])

	// membuat array secara langsung
	var number = [3]int{ 10, 11, 30 }
	// akses array menggunakan for range
	for key, _ := range number {
		fmt.Println(number[key])
	}
	// akses array menggunakan for loop
	for i := 0; i < len(number); i++  {
		fmt.Println(number[i])
	}
}
