package main

import "fmt"

func getReturnNamed() (firstName, middleName, lastName string) {
	firstName = "Chayra"
	middleName = "Ainin"
	lastName = "Zunaira"
	return
}

func main() {
	firstName, middleName, lastName := getReturnNamed()
	fmt.Println("Nama", firstName, middleName, lastName)
}
