package main

import "fmt"

func getFullName() (string, string) {
	return "Agung", "Nugraha"
}

func getIgnoreReturn() (string, string) {
	return "Hahahahah", "Wkwkwkwk"
}

func main() {
	firstname, lastname := getFullName()
	fmt.Printf("Hello my name is %s %s", firstname, lastname)

	fmt.Printf("\n")

	hahaha, _ := getIgnoreReturn()
	fmt.Printf("Ketawa %s", hahaha)
}
