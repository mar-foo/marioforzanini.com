package main

import "fmt"

func main() {
	var secondi int
	fmt.Println("Numero di secondi: ")
	fmt.Scan(&secondi)
	fmt.Println("Numero di minuti: ", secondi/60)
	return
}
