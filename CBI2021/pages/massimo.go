package main

import "fmt"

func main() {
	var somma, massimo, numero int
	fmt.Print("Inserire 5 valori")
	for i := 0; i < 4; i++ {
		fmt.Scan(&numero)
		if numero > massimo {
			massimo = numero
		}
		somma += numero
	}
	fmt.Println("Somma: ", somma)
	fmt.Println("Massimo: ", massimo)
	return
}
