package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 3
		c = a % b
	)
	fmt.Println(c)
	fmt.Println("**************")

	var i = 10
	i += 10
	i++
	fmt.Println(i)
}
