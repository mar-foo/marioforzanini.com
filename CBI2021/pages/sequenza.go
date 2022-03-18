package main

import "fmt"

func main() {
	var end int
	fmt.Print("Inserire un valore: ")
	fmt.Scan(&end)
	for i := 1; i <= end; i++ {
		fmt.Println(i)
	}
	return
}
