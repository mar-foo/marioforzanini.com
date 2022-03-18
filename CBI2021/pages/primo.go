package main

import "fmt"

func main() {
	var numero int
	for i := 2; i <= numero/2; i++ {
		if numero%i == 0 {
			fmt.Println(numero, " non è primo.")
			return
		}
	}
	fmt.Println(numero, " è primo.")
	return
}
