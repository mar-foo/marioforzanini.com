package main

import "fmt"

func main() {
	var base, altezza int
	fmt.Print("Base: ")
	fmt.Scan(&base)

	fmt.Print("Altezza: ")
	fmt.Scan(&altezza)

	rettangolo(base, altezza)
	return
}

func rettangolo(base, altezza int) {
	for i := 0; i < altezza; i++ {
		for j := 0; j < base; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
