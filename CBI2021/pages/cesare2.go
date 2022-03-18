package main

import "fmt"

func main() {
	var key int
	var c byte

	fmt.Print("Chiave: ")
	fmt.Scan(&key)
	fmt.Print("testo: ")
	for fmt.Scanf("%c", &c); c != '\n'; fmt.Scanf("%c", &c) {
		fmt.Print(string(caesar(c, key)))
	}
	fmt.Println()
	return
}

func caesar(c byte, key int) byte {
	return byte('a' + (int(c)+key-'a')%('z'-'a'+1))
}
