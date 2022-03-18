package main

import "fmt"

func main() {
	var speed, s0, t0, t1 float64

	fmt.Print("Velocità: ")
	fmt.Scan(&speed)

	fmt.Print("s0: ")
	fmt.Scan(&s0)

	fmt.Print("t0: ")
	fmt.Scan(&t0)

	fmt.Print("t1: ")
	fmt.Scan(&t1)

	fmt.Println("t\ts")
	for t := t0; t <= t1; t++ {
		fmt.Printf("%d\t%d", t, mru(speed, s0, t))
	}
	return
}

func mru(speed, s0, t float64) float64 {
	return s0 + speed * t
}
