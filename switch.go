package main

import "fmt"

func main() {
	name := "Agung Nugraha"

	switch name {
	case "Agung Nugraha":
		fmt.Println("Benar")
	default:
		fmt.Println("Ini default")
	}
}
