package main

import "fmt"

func main() {
	var gg1, start1, end1 int
	var gg2, start2, end2 int

	fmt.Println("appuntamento 1 (gg, start, end): ")
	fmt.Scan(&gg1, &start1, &end1)

	fmt.Println("appuntamento 2 (gg, start, end): ")
	fmt.Scan(&gg2, &start2, &end2)

	if gg1 != gg2 {
		fmt.Println("non si sovrappongono")
	} else {
		if !(start2 < end1 && start1 < end2) {
			fmt.Println("non si sovrappongono")
		} else {
			fmt.Println("si sovrappongono")
		}
	}
	return
}
