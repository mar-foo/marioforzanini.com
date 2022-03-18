package main

import "fmt"

func main() {
	const SECONDI_IN_ORE = 3600 // 60 * 60
	const METRI_IN_CHILOMETRO = 1000

	var chilometri, ore float64

	fmt.Println("Lunghezza del percorso (in km): ")
	fmt.Scan(&chilometri)

	fmt.Println("Tempo trascorso (in h): ")
	fmt.Scan(&ore)

	fmt.Println("Velocità media: ", chilometri/ore, "km/h")

	secondi := ore * SECONDI_IN_ORE
	metri := chilometri * METRI_IN_CHILOMETRO
	fmt.Println("Velocità media: ", metri/secondi, "m/s")

	return
}
