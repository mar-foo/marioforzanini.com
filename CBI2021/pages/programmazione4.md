# Correzione

Per prima cosa osserviamo le soluzioni agli esercizi precedenti:
[Lezione Precedente](programmazione3.html)

# Ripasso

Nella scorsa lezione abbiamo parlato di **if** ed **if/else**: le parole
chiave che forniscono un modo di cambiare il flusso di esecuzione di
un programma sulla base del verificarsi o meno di una condizione.


Il costrutto per eseguire delle istruzioni sulla base del
verificarsi della condizione **<cond>** è:


	if <cond> {
		// eseguito se <cond> == true
	} else {
		// eseguito se <cond> == false
	}

# Ripetizione di istruzioni: cicli

Molto spesso occorre ripetere la stessa sequenza di istruzioni per
ottenere il risultato desiderato, per esempio per calcolare la media
di 10 numeri occorre sommare tutti e dieci i valori e dividere il
risultato per 10; ovviamente siamo in grado di risolvere questo
problema con gli strumenti che abbiamo imparato fino ad ora:


	package main
	
	import "fmt"
	
	func main() {
	var numero1, numero2, numero3, numero4, numero5,
		numero6, numero7, numero8, numero9, numero10 float64
		var media float64
		
		fmt.Scan(&numero1, &numero2, &numero3, &numero4, &numero5,
		&numero6, &numero7, &numero8, &numero9, &numero10)
		media = (numero1 + numero2 + numero3 + numero4 + numero5 +
		numero6 + numero7 + numero8 + numero9 + numero10) / 10
		fmt.Println("Media: ", media)
		return
	}

Questa soluzione però è poco robusta: ogni volta che cambia il
numero di valori di cui fare la media occorre cambiare il programma,
ed è noiosa da scrivere: bisogna scrivere troppo. Un altro modo, più
elegante per risolvere il problema è quello di salvare i valori in
una sola variabile e, prima di leggere il valore successivo,
aggiungere il risultato alla media.


	package main
	
	import "fmt"
	
	func main() {
		var numero, media float64
		
		// Numero 1
		fmt.Scan(&numero)
		media += numero
		
		// Numero 2
		fmt.Scan(&numero)
		media += numero
		
		// Numero 3
		fmt.Scan(&numero)
		media += numero
		
		// Numero 4
		fmt.Scan(&numero)
		media += numero
		
		// Numero 5
		fmt.Scan(&numero)
		media += numero
		
		// Numero 6
		fmt.Scan(&numero)
		media += numero
		
		// Numero 7
		fmt.Scan(&numero)
		media += numero
		
		// Numero 8
		fmt.Scan(&numero)
		media += numero
		
		// Numero 9
		fmt.Scan(&numero)
		media += numero
		
		// Numero 10
		fmt.Scan(&numero)
		media += numero
		
		media /= 10
		fmt.Println("Media: ", media)
		return
	}

Questa soluzione risulta più robusta della precedente: non occorre
aggiungere una nuova variabile ogni volta che cambia il numero di
addendi, ma è comunque molto lunga e noiosa da scrivere: bisogna
ricopiare lo stesso pezzo di codice 10 volte per ottenere il
risultato.

## Ciclo for

Il primo modo di specificare un ciclo è il **for unario**: dopo la
parola chiave for occorre aggiungere un solo argomento, ovvero la
condizione da verificare ogni volta che si arriva alla fine del
blocco di istruzioni. Se la condizione specificata è vera si
ripete l'intero blocco.

	for <cond> {
		// Istruzioni che si ripetono finchè <cond> non risulta falsa
	}


*NOTA*: Negli altri linguaggi questo tipo di ciclo viene
solitamente indicato con la parola **while**.

## Ciclo for zerario

Per esprimere un ciclo che continua a ripetersi senza controllare
nessuna condizione è sufficiente non inserire nessuna condizione
dopo la parola chiave **for**


	for {
		// Istruzioni che si ripetono sempre
	}

### Break e continue

All'interno del blocco di istruzioni di un ciclo for sono
disponibili due parole chiave che permettono di controllare il
flusso di esecuzione del ciclo: **break** e **continue**. La parola
chiave break permette di uscire dal ciclo e continuare
l'esecuzione del programma a partire dalla prima istruzione
successiva. La parola chiave **continue** permette di passare
immediatamente all'iterazione successiva senza eseguire
nessun'altra istruzione.

### Esempio

Si potrebbe utilizzare un ciclo for unario al posto del costrutto
**if**:


	for <cond> {
		// Istruzioni da eseguire
		break // Esci dal ciclo
	}

## Ciclo for ternario

Il secondo modo per specificare un ciclo è il **for ternario**: dopo
la parola chiave for occorre aggiungere 3 argomenti separati da
``;``. Il primo argomento è l'istruzione che va eseguita prima di
iniziare il ciclo, il secondo argomento la condizione che viene
verificata alla fine di ogni iterazione, il terzo argomento è
l'istruzione che viene ripetuta all'inizio di ogni nuovo blocco.
Per esempio per stampare una sequenza di numeri da 1 a 10 si può
scrivere:


	for i := 1; i <= 10; i++ {
	fmt.Println(i)
	}

In questo caso l'esecuzione del codice ha questo ordine:

0. dichiaro i = 1
0. stampo i
0. se i <= 10 incremento di 1 il valore di i (è il significato di **i++**)
0. torno al passo 2

### Media

A questo punto siamo in grado di risolvere in maniera molto più
concisa il problema della media di 10 numeri che abbiamo posto
all'inizio:

	package main
	
	import "fmt"
	
	const nAddendi = 10
	func main() {
		var numero, media float64
		for i := 0; i < nAddendi; i++ {
			fmt.Scan(&numero)
			media += numero
		}
		media /= nAddendi
		fmt.Println("Media: ", media)
		return
	}

Questa risoluzione è robusta quanto quella di prima: ogni volta
che cambia il numero di addendi è sufficiente cambiare il valore
10 nel codice, ma risulta molto più breve e semplice da scrivere.
Resta il problema di non essere in grado di calcolare la media in
base al numero di addendi inseriti dall'utente (senza fissare il
numero a priori).

## Variabili locali

Come per i costrutti di selezione **if** e **if/else** le variabili
che vengono dichiarate negli argomenti del **for** sono locali.

# Esercizi

## Esercizio 0 - Raddoppio

*Problema:* Scrivere un programma go ``raddoppio.go`` che legge 5
numeri da tastiera e di ciascuno stampa il doppio.


*Esempi di esecuzione*:

Inserire un valore: **2**

4

Inserire un valore: **5**

10

Inserire un valore: **7**

14

Inserire un valore: **3**

6

Inserire un valore: **53**

106

## Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var numero int
		for i := 0; i < 4; i++ {
			fmt.Print("Inserire un valore: ")
			fmt.Scan(&numero)
			fmt.Println(numero * 2)
		}
		return
	}

## Esercizio 1 - Sequenza

*Problema:*  Scrivere un programma go ``sequenza.go`` che legge un
numero da tastiera e stampa la sequenza di valori da 1 a n.


*Esempio di esecuzione*:

Inserire un valore: **5**

1

2

3

4

5

### Soluzione
	
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

## Esercizio 2 - Sequenza pari

*Problema:* Scrivere un programma go ``sequenza_pari.go`` che legge
un numero da tastiera e stampa la sequenza di valori pari tra 1 e
n.

*Esempio di esecuzione*:

Inserire un valore: **5**

2

4

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var numero int
		fmt.Print("Inserire un valore: ")
		fmt.Scan(&numero)
		for i := 2; i <= numero; i = i + 2 {
			fmt.Println(i)
		}
		return
	}

## Esercizio 3 - Tabellina

*Problema:* Scrivere un programma go ``tabellina.go`` che legge un
numero da tastiera e stampa la tabellina di quel numero (n*1, n*2,
&#x2026;, n*10).

*Esempio di esecuzione*:

Inserire un valore: **3**

3

6

9

12

15

18

21

24

27

30

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var numero int
		fmt.Print("Inserire un valore: ")
		fmt.Scan(&numero)
		
		for i := 1; i <= 10; i++ {
			fmt.Println(i * numero)
		}
		return
	}

## Esercizio 4 - Voto valido

*Problema:* Scrivere un programma go ``voto_valido.go`` che stampa
"voto:" per chiedere un valore voto fino ad ottenere un valore
valido, cioè compreso tra 3 <= voto <= 10, e poi stampa "voto
valido"



*Esempio di esecuzione*:

Inserire un voto: **11**

voto non valido

Inserire un voto: **2**

voto non valido

Inserire un voto: **5**

voto valido

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var voto int
		fmt.Print("Inserire un voto: ")
		fmt.Scan(&voto)
		for voto < 3 || voto > 10 {
			fmt.Print("voto non valido")
			fmt.Print("Inserire un voto: ")
			fmt.Scan("&voto")
		}
		fmt.Println("voto valido")
		return
	}

## Esercizio 5 - Massimo

*Problema:* Scrivere un programma ``massimo.go`` che legge 5 numeri
numeri interi, ne stampa la somma ed il massimo.

*Esempio di esecuzione*:

Inserire 5 valori: **1 3 2 5 4**

Somma: 15

Massimo: 5

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var somma, massimo, numero int
		fmt.Print("Inserire 5 valori")
		for i := 0; i < 4; i++ {
			fmt.Scan(&numero)
			if numero > massimo {
				massimo = numero
			}
			somma += numero
		}
		fmt.Println("Somma: ", somma)
		fmt.Println("Massimo: ", massimo)
		return
	}

## Esercizio 6 - Primo

*Problema:* Scrivere un programma ``primo.go`` che legge in input un
numero intero e determina se è primo.

*Esempio di esecuzione*:

Inserire un valore: **7**

È primo

*Nota*: Per stabilire se un valore è primo è sufficiente
determinare il primo numero che è suo divisore (se c'è).

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var numero int
		for i := 2; i <= numero / 2; i++ {
			if numero % i == 0 {
				fmt.Println(numero, " non è primo.")
			return
			}
		}
		fmt.Println(numero, " è primo.")
		return
	}

## Esercizio 7 - Media

*Problema:* Scrivere un programma ``media_var.go`` che legge in input
dei numeri diversi da 0 e ne stampa la media.

*Esempio di esecuzione*:

Inserire un valore: **5**

Inserire un valore: **4**

Inserire un valore: **6**

Inserire un valore: **0**

Media: 5

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var valore, quantità, media float64
		fmt.Print("inserire un valore: ")
		fmt.Scan(&valore)
		for numero != 0 {
			media += valore
			quantità++
			fmt.Print("inserire un valore: ")
			fmt.Scan(&valore)
		}
		media /= quantità
		fmt.Println("Media: ", media)
		return
	}
