package main

import (
	"fmt"
	"unicode"
)

func main() {
	var c rune
	fmt.Print("Inserire una stringa: ")
	first := true
	for fmt.Scanf("%c", &c); c != '\n'; fmt.Scanf("%c", &c) {
		if first && ! unicode.IsSpace(c) {
			first = false
			fmt.Print("%c", unicode.ToUpper(c))
			continue
		} else if unicode.IsSpace(c) {
			first = true
		}
		fmt.Printf("%c", unicode.ToLower(c))
	}
	fmt.Println()
	return
}
