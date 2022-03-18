package main

import "fmt"

func main() {
	var n int

	// Punto a
	fmt.Println("int uguale a 10: ")
	fmt.Scan(&n)
	fmt.Println(n == 10)

	// Punto b
	fmt.Println("int diverso da 10: ")
	fmt.Scan(&n)
	fmt.Println(n != 10)

	// Punto c
	fmt.Println("int diverso da 10 e da 20: ")
	fmt.Scan(&n)
	fmt.Println(n != 10 && n != 20)

	// Punto d
	fmt.Println("int diverso da 10 o da 20: ")
	fmt.Scan(&n)
	fmt.Println(n != 10 || n != 20)

	// Punto e
	fmt.Println("int maggiore o uguale a 10: ")
	fmt.Scan(&n)
	fmt.Println(n >= 10)

	// Punto f
	fmt.Println("int compreso tra 10 e 20, inclusi: ")
	fmt.Scan(&n)
	fmt.Println(n <= 20 && n >= 10)

	// Punto g
	fmt.Println("int compreso tra 10 e 20, esclusi: ")
	fmt.Scan(&n)
	fmt.Println(n < 20 && n > 10)

	// Punto h
	fmt.Println("int compreso tra 10 e 20, 10 incluso e 20 escluso: ")
	fmt.Scan(&n)
	fmt.Println(n < 20 && n >= 10)

	// Punto i
	fmt.Println("int minore di 10 o maggiore di 20")
	fmt.Scan(&n)
	fmt.Println(n < 10 || n > 20)

	// Punto j
	fmt.Println("int compreso tra 10 e 20 o tra 30 e 40, estremi inclusi: ")
	fmt.Scan(&n)
	fmt.Println((n <= 20 && n >= 10) || (n <= 40 && n >= 30))

	// Punto k
	fmt.Println("int multiplo di 4 ma non di 100")
	fmt.Scan(&n)
	fmt.Println(n%4 == 0 && n%100 != 0)

	// Punto l
	fmt.Println("int dispari e compreso tra 0 e 100, inclusi")
	fmt.Scan(&n)
	fmt.Println(n%2 == 1 && 0 <= n && a <= 100)

	return
}
