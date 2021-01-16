package main

import "fmt"

func main() {
	var (
		nilai32 int32 = 100000
		nilai64 int64 = int64(nilai32)

		// pastikan int sesuai dengan besar nilai yg dapat ditampung
		// jika melebihin dari kapasitas penyimpanan
		// maka akan mengembalikan data yg salah
		nilai8 int8 = int8(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// conversion type data string ke byte
	// kemudian mengembalikannya lagi
	var (
		name = "Agung Nugraha"
		e byte = name[0]
		eString string = string(e)
	)

	fmt.Println(name)
	fmt.Println(eString)

}
