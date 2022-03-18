package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)

	maiuscole, minuscole := false, false
	for _, c := range {
		if unicode.IsUpper(c) {
			maiuscole = true
		} else if unicode.IsLower(c) {
			minuscole = true
		}
	}

	if minuscole && maiuscole {
		fmt.Println("Sia maiuscole che minuscole")
	} else if maiuscole {
		fmt.Println("Solo maiuscole")
	} else if minuscole {
		fmt.Println("Solo minuscole")
	}
	return
}
