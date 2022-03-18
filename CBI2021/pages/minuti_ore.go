package main

import "fmt"

func main() {
	var ore int
	fmt.Println("Numero di ore:")
	fmt.Scan(&ore)
	fmt.Println("Numero di giorni: ", ore/24)
	return
}
