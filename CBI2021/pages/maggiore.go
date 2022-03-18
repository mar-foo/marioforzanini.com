package main

import "fmt"

func main() {
	var max, min int

	fmt.Println("due int:")
	fmt.Scan(&max)
	fmt.Scan(&min)

	if min > max {
		max, min = min, max // Scambio i valori usando l'assegnamento multiplo
	}

	fmt.Println(max, min)

	return
}
