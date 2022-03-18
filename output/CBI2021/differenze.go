package main

import "fmt"

func main() {
	var corrente, precedente float64
	fmt.Print("Scrivi un numero: ")
	fmt.Scan(&corrente)
	for corrente != 0 {
		precedente = corrente
		fmt.Print("Scrivi un numero: ")
		fmt.Scan(&corrente)
		fmt.Println("Differenza = ", corrente-precedente)
	}
	return
}
