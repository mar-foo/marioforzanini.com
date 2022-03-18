package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var c rune

	fmt.Print("Carattere da cercare: ")
	fmt.Scanf("%c", &c)

	fmt.Print("Stringa: ")
	fmt.Scan(&s)

	fmt.Println(strings.IndexRune(s, c))
	return
}
