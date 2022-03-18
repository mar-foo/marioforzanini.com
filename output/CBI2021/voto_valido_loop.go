package main

import "fmt"

func main() {
	var voto int
	fmt.Print("Inserire un voto: ")
	fmt.Scan(&voto)
	for voto < 3 || voto > 10 {
		fmt.Print("voto non valido")
		fmt.Print("Inserire un voto: ")
		fmt.Scan("&voto")
	}
	fmt.Println("voto valido")
	return
}
