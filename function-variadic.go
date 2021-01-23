package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	sum := sumAll(10, 30, 1, 21, 25)
	fmt.Println("Hasil :", sum)

	slice := []int{10, 30, 20, 51, 21}
	total := sumAll(slice...)
	fmt.Println("Hasilnya :", total)
}
