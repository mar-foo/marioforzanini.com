package main

import "fmt"

func main() {
	const ORE_IN_GIORNI = 24
	const MINUTI_IN_ORE = 60
	const SECONDI_IN_MINUTO = 60

	var secondiTotali int

	fmt.Println("Numero di secondi:")
	fmt.Scan(&secondiTotali)

	secondi_in_ore := MINUTI_IN_ORE * SECONDI_IN_MINUTO
	secondi_in_giorni := ORE_IN_GIORNI * secondi_in_ore

	giorni := secondiTotali / secondi_in_giorni
	resto := secondiTotali % secondi_in_giorni

	ore := resto / secondi_in_ore
	resto = resto % secondi_in_ore

	minuti := resto / SECONDI_IN_MINUTO
	secondi := resto % SECONDI_IN_MINUTO

	fmt.Println("g:h:m:s -", giorni, ":", ore, ":", minuti, ":", secondi)
	return
}
