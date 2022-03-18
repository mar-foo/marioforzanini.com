package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Inserire una parola: ")
	fmt.Scan(&s)

	fmt.Println(contaVocali(s))
	return
}

func contaVocali(s string) int {
	var count int

	count += strings.Count(s, "a")
	count += strings.Count(s, "e")
	count += strings.Count(s, "i")
	count += strings.Count(s, "o")
	count += strings.Count(s, "u")
	count += strings.Count(s, "A")
	count += strings.Count(s, "E")
	count += strings.Count(s, "I")
	count += strings.Count(s, "O")
	count += strings.Count(s, "U")

	return count
}
