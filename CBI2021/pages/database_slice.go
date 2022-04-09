package main

import (
	"fmt"
	"io"
	"os"
)

var matematica, fisica, informatica []float64

func main() {
	err := leggiVoti("database.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Medie:")
	fmt.Println("Matematica: ", media(matematica))
	fmt.Println("Fisica: ", media(fisica))
	fmt.Println("Informatica: ", media(informatica))

	return
}

func leggiVoti(filename string) error {
	var materia string
	matematica = make([]float64, 0)
	fisica = make([]float64, 0)
	informatica = make([]float64, 0)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	for err != io.EOF {
		var voto float64
		if _, err = fmt.Fscanf(f, "%s %f\n", &materia, &voto); err != nil && err != io.EOF {
			return err
		}

		switch materia {
		case "matematica":
			matematica = append(matematica, voto)
		case "fisica":
			fisica = append(fisica, voto)
		case "informatica":
			informatica = append(informatica, voto)
		}
	}

	f.Close()
	return nil
}

// Restituisce la media dei voti
func media(voti []float64) float64 {
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

	return sum / float64(len(voti))
}
