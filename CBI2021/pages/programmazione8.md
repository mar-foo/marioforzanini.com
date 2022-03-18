# Ripasso

Nelle scorse lezioni ci siamo concentrati su 3 concetti:


0. Variabili
0. if/else
0. Cicli for

Questi tre componenti sono sufficienti per scrivere qualsiasi
programma Go, ma non rappresentano tutto l'arsenale messo a
disposizione dal linguaggio.  Scrivere programmi complessi con questi
strumenti è scomodo e sconsigliato: occorre la possibilità di
organizzare meglio il codice.

# Pacchetti

I **pacchetti** sono una maniera di
strutturare il codice: un programma è
costruito come un pacchetto e può usare
strumenti messi a disposizione da altri
pacchetti.


Ogni file appartiene ad un solo pacchetto,
indicato alla prima riga con la parola chiave
**package**, ed ogni pacchetto può essere
composto da più file.Tutte le applicazioni Go
contengono un pacchetto chiamato
**main**.


Un programma Go viene costruito collegando
un set di pacchetti attraverso la parola
chiave **import**.  Quando nei nostri
programmi scrivevamo


	import "fmt"

Stavamo indicando che il nostro programma
aveva bisogno degli strumenti messi a
disposizione del pacchetto *"fmt"* (che
contiene funzionalità di Input/Output).  Ogni
pacchetto può essere richiamato in uno
statement separato:


	import "fmt"
	import "os"

Oppure in una maniera più sintetica:

	import (
		"fmt"
		"os"
	)


Go mette a disposizione una serie di **pacchetti** (in altri linguaggi
sono anche chiamate *librerie*) che implementano funzionalità utili:
il protocollo HTTP (per i server web), lettura e scrittura da file,
operazioni sulle stringhe (capitalizzazione, ricerca di
sottostringhe...), funzioni matematiche (seno, coseno, potenze,
radici...) ecc.


Dopo gli **import** si possono dichiarare delle costanti (**const**) o
variabili (**var**), questi identificatori sono *globali*, ovvero sono
disponibili in tutto il pacchetto.

## Visibilità

Quando un identificatore (variabile, costante...) inizia con la
lettera **maiuscola**, l'oggetto a cui si riferisce è visibile fuori
dal pacchetto (viene **esportato**). Gli identificatori che iniziano
con la lettera minuscola sono visibili soltanto nel pacchetto in cui
sono definiti.

## Accedere agli oggetti esportati

Supponendo di aver importato un pacchetto ``pkg`` che contiene un
simbolo ``Simbolo``, ci si può riferire a ``Simbolo`` con la notazione
``pkg.Simbolo``. Questo permette di non avere conflitti tra
identificatori uguali definiti in pacchetti differenti.


# Funzioni

Nei nostri programmi abbiamo già avuto modo di utilizzare delle
**funzioni**, per esempio ``fmt.Print`` oppure ``fmt.Scan`` (entrambe
dal pacchetto fmt).

In Go una funzione è un oggetto che accetta degli argomenti in input e
usa quegli argomenti per fare qualcosa.

Questa definizione è volutamente vaga, poichè lo spettro di azioni
possibili con una funzione è molto ampio:

0. Stampare oggetti (``fmt.Print``)
0. Leggere valori ed inserirli in una (o più) variabili (``fmt.Scan``)
0. Leggere un numero e restituirne la potenza di 10 corrispondente (``math.Pow10``)

Inoltre una funzione può restituire dei valori in output, per esempio
``math.Pow10`` accetta in input un valore ``float64`` e restituisce un
valore ``float64``.

La parola chiave per dichiarare una nuova funzione è **func** e la sintassi è la seguente:

	func nomeFunzione(lista_parametri) (lista_valori_output) {
		// Codice
	}

Dove

* *nomeFunzione* è l'identificatore che servirà a richiamare la funzione (come ``Print`` per ``fmt.Print``)
* *lista_parametri* è la lista di valori che la funzione accetta in input con i loro tipi
* *lista_valori_output* è la lista di valori che la funzione restituisce come output, con i loro tipi

## Esempio

Una funzione che calcola la media di due numeri può essere scritta così:

	func media(primo float64, secondo float64) float64 {
		return (primo + secondo) / 2
	}

La parola chiave **return** ferma l'esecuzione della funzione dando in
output ``(primo + secondo) / 2``.

Questa funzione può essere usata per scrivere un programma che calcola
la media di due numeri letti dall'utente:

	package main
	import "fmt"
	
	func main() {
		var x, y float64
		fmt.Print("Inserire due numeri: ")
		fmt.Scan(&x, &y)
		fmt.Println("Media: ", media(x, y))
		return
	}
	
	func media(primo float64, secondo float64) float64 {
		return (primo + secondo) / 2
	}


## NB:

I nomi dei parametri ``primo`` e ``secondo`` non c'entrano nulla con i
nomi delle variabili che vengono ricevute in input: quando la funzione
viene chiamata, per esempio nella funzione `` main()``, i valori di
input vengono semplicemente copiati dentro due variabili locali alla
funzione chiamate ``primo`` e ``secondo`` ``.

# Documentazione

Go fornisce uno strumento molto utile per navigare nella
documentazione dei pacchetti standard, si tratta del programma ``go
doc`` che accetta come argomento il nome di un simbolo o di un
pacchetto e ne stampa la documentazione. Ad esempio per vedere la
documentazione di ``fmt.Println``:

	go doc fmt.Println

Oppure per vedere la documentazione del pacchetto ``fmt`` (compresa la
lista di simboli esportati dal pacchetto):

	go doc fmt

# Esercizi

## Esercizio 0 - Euclide

Riscrivete l'[Esercizio 2 della scorsa lezione](programmazione6.html) utilizzando
una funzione ``euclide`` che accetti in input due numeri interi e ne restituisca l'MCD utilizzando l'algoritmo di Euclide.

## Esercizio 1 - Slash

Riscrivete l'[Esercizio 5 della scorsa lezione](programmazione6.html) utilizzando
una funzione "slash". Quali argomenti avrà questa funzione? Quali valori restituisce?

## Esercizio 2 - Rettangolo

*Problema*: Scrivere un programma ``rettangolo.go`` che legge due interi positivi ``n`` e
``m`` e stampa un rettangolo di base ``n`` e altezza ``m`` asterischi. L'implementazione deve
contenere una funzione chiamata "rettangolo" che si occupi di disegnare la  figura.

*Esempi di esecuzione*:

Base: **5**

Altezza: **4**

	*****
	*****
	*****
	*****


## Esercizio 3 - Legge oraria

*Problema*: Scrivere un programma ``legge_oraria.go`` che accetti in input 4 parametri:

* Velocità
* Posizione al tempo 0
* Tempo iniziale t₀
* Tempo finale t₁

e stampi una tabella delle posizioni del corpo soggetto alla legge oraria:

	s = s₀ + v * t

di secondo in secondo nell'intervallo ``t₀ ≤ t ≤ t₁``

*Esempi di esecuzione*:

Velocità: **5**

s₀: **1**

t₀: **10**

t₁: **20**

	t   s
	10  51
	11  56
	12  61
	13  66
	14  71
	15  76
	16  81
	17  86
	18  91
	19  96
	20  101


## Esercizio 4 - Piano inclinato

*Problema*: Scrivere un programma ``piano_inclinato.go`` che legga 2 parametri:

* Massa ``m``
* Angolo alla base ``α``

e calcoli la forza necessaria a mantenere in equilibrio un corpo di
massa ``m`` su un piano inclinato ``α`` gradi rispetto al suolo.
Trascurate ogni attrito.

*Esempi di esecuzione*:

Massa: **14.2**

Angolo alla base: **45**

Forza necessaria: 10.04

*Nota*: Go fornisce funzioni che calcolano seno, coseno, tangente ... di angoli espressi in radianti. Per trasformare un angolo da gradi a radianti è sufficiente questa proporzione:

	π rad : 360° = x rad : α °

se ``x`` è l'angolo in radianti e α l'angolo in gradi. Perciò:

	x = (α * π) / 360

il valore π si trova nella costante ``math.Pi``.

*Consiglio*: Provate il comando ``go doc math``