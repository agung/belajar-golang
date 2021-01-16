package main

import "fmt"

// membuat alias untuk type data
func main() {
	type myString string

	var name myString = "Agung Nugraha"
	fmt.Println(name)
}
