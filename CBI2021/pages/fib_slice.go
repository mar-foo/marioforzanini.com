package main

import (
	"fmt"
	"os"
)

func main() {
	var max int
	fmt.Print("Massimo numero della sequenza: ")
	for _, err := fmt.Scan(&max);  err != nil; {
		fmt.Fprintln(os.Stderr, err)
		_, err = fmt.Scan(&max)
	}

	fmt.Println(fib(max))
}

func fib(n int) []int {
	var fibonacci []int

	fibonacci = make([]int, n)

	if n > 1 {
		fibonacci[0] = 0
		fibonacci[1] = 1
		for i := 1; i < len(fibonacci); i++ {
			fibonacci[i + 1] = fibonacci[i] + fibonacci[i - 1]
		}
	}
	return fibonacci
}
