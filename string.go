package main

import "fmt"

func main() {
	// concat string
	firstName := "Agung"
	lastName := "Nugraha"

	fmt.Println(firstName + " " + lastName) // Print nama as string
	full := firstName + " " + lastName

	// count character string
	fmt.Println(len(firstName + " " + lastName))
	
	// loop character string
	for _, value := range full {
		fmt.Println(string(value))
	}
}
