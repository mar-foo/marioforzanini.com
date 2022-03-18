package main

import "fmt"

func main() {
	var distanza, carburante float64
	fmt.Println("Distanza percorsa (in km): ")
	fmt.Scan(&distanza)
	fmt.Println("Quantità di carburante utilizzata (in l): ")
	fmt.Scan(&carburante)

	fmt.Println("Consumo medio: ", carburante/distanza, "km/l")
	fmt.Println("Resa media: ", distanza/carburante, "l/km")
}
