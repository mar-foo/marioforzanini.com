package main

import (
	"fmt"
	"os"
)

func main() {
	voti, err := leggiVoti("voti.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // Qualunque numero diverso da zero indica un errore
	}

	fmt.Println("Media: ", media(voti))
	return
}

// Legge un array di 10 voti dal file filename
func leggiVoti(filename string) ([10]float64, error) {
	var voti [10]float64
	f, err := os.Open(filename)
	if err != nil {
		return voti, err
	}

	for i := 0; i < 10; i++ {
		_, err = fmt.Fscan(f, &voti[i])
		if err != nil {
			return voti, err
		}
	}

	f.Close()
	return voti, nil
}

// Restituisce la media dei voti
func media(voti [10]float64) float64 {
	var sum float64

	for i := range voti {
		if voti[i] < 3 {
			fmt.Fprintf(os.Stderr, "Voto non valido %f, utilizzo il valore valido più vicino: %d\n", voti[i], 3)
			sum += 3
		} else if voti[i] > 10 {
			fmt.Fprintf(os.Stderr, "Voto non valido %f, utilizzo il valore valido più vicino: %d\n", voti[i], 10)
			sum += 10
		} else {
			sum += voti[i]
		}
	}

	return sum / 10
}
