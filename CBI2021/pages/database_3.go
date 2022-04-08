package main

import (
	"fmt"
	"os"
)

var matematica, fisica, informatica [10]float64

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
	var matIndex, fisIndex, infIndex int

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	for matIndex <= 9 || infIndex <= 9 || fisIndex <= 9 {
		var voto float64
		if _, err = fmt.Fscanf(f, "%s %f\n", &materia, &voto); err != nil {
			return err
		}

		switch materia {
		case "matematica":
			if matIndex < 10 {
				matematica[matIndex] = voto
			} else {
				fmt.Fprintf(os.Stderr, "Numero massimo di voti di %s raggiunto, ignoro %f\n", materia, voto)
			}
			matIndex++
		case "fisica":
			if fisIndex < 10 {
				fisica[fisIndex] = voto
			} else {
				fmt.Fprintf(os.Stderr, "Numero massimo di voti di %s raggiunto, ignoro %f\n", materia, voto)
			}
			fisIndex++
		case "informatica":
			if infIndex < 10 {
				informatica[infIndex] = voto
			} else {
				fmt.Fprintf(os.Stderr, "Numero massimo di voti di %s raggiunto, ignoro %f\n", materia, voto)
			}
			infIndex++
		}
	}

	f.Close()
	return nil
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
