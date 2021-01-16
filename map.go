package main

import "fmt"

// map adalah type data array dengan menggunakan
// key berupa string, seperti array associative pada php
func main() {
	person := map[string]string{
		"name": "Agung",
		"age": "24",
	}

	fmt.Println(person["name"])

	book := make(map[string]string)
	book["name"] = "Agung Nugraha"
	book["age"] = "24"
	fmt.Println(book)

	delete(book, "age")
	fmt.Println(book)

}
