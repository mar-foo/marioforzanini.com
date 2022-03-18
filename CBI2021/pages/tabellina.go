package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Inserire un valore: ")
	fmt.Scan(&numero)

	for i := 1; i <= 10; i++ {
		fmt.Println(i * numero)
	}
	return
}
