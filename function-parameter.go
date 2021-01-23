package main

import "fmt"

// function with parameter
func sayHello(firstName string, lastName string) {
	fmt.Printf("Hello %s %s", firstName, lastName)
}

func main() {
	firstName := "Agung"
	lastName := "Nugraha"

	sayHello(firstName, lastName)
}
