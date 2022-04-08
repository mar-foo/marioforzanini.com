package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib())
	return
}

func fib() [10]int {
	var fibonacci [10]int

	fibonacci[0] = 0
	fibonacci[1] = 1

	for j := 1; j < 9; j++ {
		fibonacci[j + 1] = fibonacci[j] + fibonacci[j - 1]
	}
	return fibonacci
}
