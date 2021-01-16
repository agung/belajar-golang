package main

import (
	"fmt"
	"strconv"
)

// deklarasi single constanta
const basic string  = "Basic"

// deklarasi multiple constanta
const (
	name = "Agung Nugraha"
	age = 24
)

func getName() {
	say := "Hello my name is " + name + ", i am " + strconv.Itoa(age) + " years old"
	fmt.Println(say)
}

func main() {
	getName()

	// memanggil constant
	fmt.Println(basic)
}
