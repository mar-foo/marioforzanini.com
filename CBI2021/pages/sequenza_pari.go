package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Inserire un valore: ")
	fmt.Scan(&numero)
	for i := 2; i <= numero; i = i + 2 {
		fmt.Println(i)
	}
	return
}
