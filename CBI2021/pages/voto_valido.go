package main

import "fmt"

func main() {
	var voto int
	fmt.Println("voto: ")
	fmt.Scan(&voto)

	if voto < 3 || voto > 10 {
		fmt.Println("voto non valido")
	}

	return
}
