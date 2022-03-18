package main

import "fmt"

func main() {
	var n, k int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 1; i*i < n; i++ {
		k = i * i
	}
	fmt.Println("Massimo quadrato: ", k)
	return
}
