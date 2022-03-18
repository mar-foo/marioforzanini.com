package main

import "fmt"

func main() {
	var primoVoto, secondoVoto float64

	fmt.Println("Primo voto: ")
	fmt.Scan(&primoVoto)

	fmt.Println("Secondo voto: ")
	fmt.Scan(&secondoVoto)

	fmt.Println("Media: ", (primoVoto+secondoVoto)/2)
	return
}
