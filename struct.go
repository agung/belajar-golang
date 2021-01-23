package main

import (
	"fmt"
	"tutorial/customer"
)

func main() {
	customer.SetName("Agung Nugraha")

	fmt.Println(customer.GetName())
}
