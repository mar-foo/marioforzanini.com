package main

import "fmt"

func main() {
	var eta int
	var studente bool

	fmt.Println("età:")
	fmt.Scan(&eta)

	fmt.Println("studente? (t/f):")
	fmt.Scan(&studente)

	if eta >= 0 && eta < 9 {
		fmt.Println("gratis")
	} else if studente || eta < 14 {
		fmt.Println("5 euro")
	} else if (eta >= 14 && eta < 26) || eta > 65 {
		fmt.Println("7.5 euro")
	} else {
		fmt.Println("10 euro")
	}

	return
}
