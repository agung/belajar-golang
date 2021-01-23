package main

import "fmt"

func catch() {
	message := recover()
	if message != nil {
		fmt.Println(message)
	}
}

func sayHello(name string) string {
	if name != "nugraha" {
		panic("Error args " + name)
	}
	return name
}

func main() {
	// run function logging on final exec
	// event on panic
	defer catch()

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(sayHello("agung"))
}
