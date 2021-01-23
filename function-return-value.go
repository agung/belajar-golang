package main

import "fmt"

func sayHello(name string) string {
	return fmt.Sprintf("Hello my name is %s", name)
}

func main() {
	hello := sayHello("Agung Nugraha")
	fmt.Println(hello)
}
