package main

import "fmt"

func main() {
	var numero int
	for i := 0; i < 4; i++ {
		fmt.Print("Inserire un valore: ")
		fmt.Scan(&numero)
		fmt.Println(numero * 2)
	}
	return
}
