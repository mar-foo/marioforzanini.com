package main

import "fmt"

func main() {
	var input string
	fmt.Print("Inserire una riga del database: ")
	fmt.Scan(&input)
	database(input)
	return
}

func database(s string) {
	var index int
	var currentField string

	for _, c := range s {
		if c != ';' {
			currentField += string(c)
		} else {
			switch index {
			case 0:
				fmt.Println("Nome: ", currentField)
			case 1:
				fmt.Println("Cognome: ", currentField)
			case 2:
				fmt.Println("Email: ", currentField)
			case 3:
				fmt.Println(": ", currentField)
			}
			index++
		}
	}
}
