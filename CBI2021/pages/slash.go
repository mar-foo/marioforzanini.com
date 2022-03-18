package main

import "fmt"

func main() {
	var lines int
	fmt.Print("Altezza: ")
	fmt.Scan(&lines)
	for i := 0; i < lines; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
	return
}
