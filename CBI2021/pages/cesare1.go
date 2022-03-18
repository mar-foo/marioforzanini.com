package main

import (
	"fmt"
	"strings"
)

func main() {
	var key int
	var s string

	fmt.Print("Chiave: ")
	fmt.Scan(&key)
	fmt.Print("testo: ")
	fmt.Scan(&s)

	fmt.Println(caesar(s, key))
	return
}

func caesar(s string, key int) string {
	var result string
	s = strings.ToLower(s)
	for _, c := range s {
		result += string('a' + (int(c)+key-'a')%('z'-'a'+1))
	}
	return result
}
