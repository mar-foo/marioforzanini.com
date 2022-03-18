package main

import "fmt"

func main() {
	var c rune
	fmt.Print("Inserire un carattere: ")
	fmt.Scanf("%c", &c)

	fmt.Printf("%c%c%c", c - 1, c, c + 1)
	if 'A' < c && c < 'L' {
		fmt.Println("A-L")
	}
	return
}
