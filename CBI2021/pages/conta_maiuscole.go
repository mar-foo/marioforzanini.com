package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&input)
	fmt.Println("Maiuscole: ", contaMaiuscole(input))
	return
}

func contaMaiuscole(s string) int {
	var count int

	for _, c := range s {
		if unicode.IsUpper(c) {
			count++
		}
	}
	return count
}
