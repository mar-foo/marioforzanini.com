package main

import "fmt"

func main() {
	var valore, quantità, media float64
	fmt.Print("inserire un valore: ")
	fmt.Scan(&valore)
	for valore != 0 {
		media += valore
		quantità++
		fmt.Print("inserire un valore: ")
		fmt.Scan(&valore)
	}
	media /= quantità
	fmt.Println("Media: ", media)
	return
}
