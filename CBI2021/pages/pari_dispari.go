package main

import "fmt"

func main() {
	var numero int

	fmt.Println("numero:")
	fmt.Scan(&numero)

	if numero%2 == 0 {
		fmt.Println(numero, " è pari")
	} else {
		fmt.Println(numero, " è dispari")
	}

	return
}
