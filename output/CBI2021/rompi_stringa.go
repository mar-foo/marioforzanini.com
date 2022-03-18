package main

import "fmt"

func main() {
	var s string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&s)

	rompi(s)
	return
}

func rompi(s string) {
	for _, c := range s {
		fmt.Println(c)
	}
}
