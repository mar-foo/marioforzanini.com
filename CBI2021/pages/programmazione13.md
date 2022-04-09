# Ripasso

Settimana scorsa abbiamo parlato di gestione di errori e di array. Mi
pare che questo secondo argomento vi abbia un po' confuso, quindi
prima di concluderlo occorre riprendere alcuni concetti importanti che
forse non sono risultati chiari.

## Struttura di un array

Gli array sono uno dei tipi che Go mette a disposizione per la
gestione di una collezione di dati, ovvero di un numero di variabili
che hanno qualcosa in comune. Solitamente i dati che vanno conservati
in un array corrispondono ad una lista di dati simili (ovvero dello
stesso tipo).

Possiamo pensare ai dati di un array come ad una lista di variabili
dello stesso tipo conservate in una zona contigua di memoria. La
sintassi per dichiarare un array è la seguente:

	var nome [nelem]tipo

dove ``nome`` sarà l'identificatore con cui riferirsi alla variabile
di tipo array, ``nelem`` è la lunghezza dell'array (il numero di
elementi) e ``tipo`` è il tipo di ciascuno degli elementi (può essere
uno di quelli che abbiamo già visto: ``float64``, ``int``, ``bool``).
In particolare ``nelem`` deve essere un numero costante, non una
variabile: può essere un letterale, per esempio ``10``, oppure un
valore dichiarato con la parola chiave ``const``.

## Indicizzazione

Quello che mi pare vi abbia confuso maggiormente è l'accesso agli
elementi di un array, e di conseguenza la modifica degli elementi.

Per riferirsi a ciascun elemento dell'array si utilizza un indice,
ovvero un numero che corrisponde alla posizione dell'elemento
nell'array. Ciascun array è indicizzato a partire da **0**, ovvero il
primo elemento è lo *0-esimo*, il secondo elemento è il *primo* e così
discorrendo.

Supponiamo di avere un array di 10 interi dichiarato in questo modo:

	var list [10]int

la variabile che contiene il primo elemento dell'array è identificata
da:

	list[0]

Questo identificatore permette di utilizzare il primo elemento
dell'array in un calcolo, oppure di modificarlo: permette di
utilizzarlo come una delle variabili che abbiamo usato fino ad ora.

## Accesso a tutti gli elementi

Siccome ad un array corrisponde una lista di elementi, per accedere a
tutti i valori occorre un ciclo **for**.  Go mette a disposizione due
tipi di cicli for per lavorare su array: il primo è un semplice ciclo
che conti da **0** a **nelem - 1** ed usi l'indicizzazione che abbiamo
visto prima

	// Stampo tutti gli elementi di list
	for i := 0; i < nelem; i++ {
		fmt.Println(list[i])
	}

il secondo è un ciclo **for-range** che permette di listare gli
elementi di un array senza dover specificare la lunghezza (Go la
conosce a priori perchè l'abbiamo specificata nella dichiarazione di
``list``).

	// Stampo tutti gli elementi di list
	for i := range list {
		// la variabile i contiene un indice che va da 0 a nelem - 1
		// e cambia ad ogni iterazione
		fmt.Println(list[i])
	}

Con le operazioni di indicizzazione ed i cicli for si possono svolgere
la maggior parte delle operazioni sugli array.

# Slice

Gli array sono molto comodi, ma spesso occorre trattare i dati
indipendentemente dal loro numero, in particolare se si ottiene
l'input in maniera dinamica (chiedendolo all'utente, oppure leggendo
un file) non è sempre nota a priori la quantità di dati che bisogna
trattare. Per risolvere i problemi dovuti alla staticità della
dimensione degli array, Go mette a disposizione un altra collezione di
elementi chiamata **slice**. Una slice è in pratica un array di
dimensioni variabili, Go si preoccupa di aumentarne la grandezza
quando necessario, ed è sempre possibile aggiungere elementi alla
lista.

La sintassi per dichiarare una slice è la seguente:

	var nome []tipo

In seguito a questa dichiarazione avremo a disposizione una slice di
elementi di tipo ``tipo``, identificata da ``nome``; ciascuno degli
elementi della slice corrisponderà allo zero del tipo selezionato.

## Operazioni sulle slice

La natura dinamica delle slice richiede alcune operazioni particolari
per inizializzare ed aggiungere elementi.

## Inizializzare

Dopo una dichiarazione di questo tipo

	var list []float64

abbiamo a disposizione una slice vuota di ``float64``, se sappiamo che
la slice avrà grandezza minima 10 possiamo utilizzare la funzione
``make`` per inizializzarla:

	list = make([]float64, 10)

Una volta inizializzata possiamo accedere a tutti gli elementi tra
``list[0]`` e ``list[9]``, ma non oltre.

## Aggiungere elementi

Per aggiungere elementi in coda alla slice si usa la funzione
``append``:

	// func append(slice []Type, elems ...Type) []Type
	list = append(list, 43.22)

append restituisce una nuova slice, fatta dalla slice di partenza e
dagli altri elementi aggiunti in coda.

In particolare se non si conosce la grandezza della slice si può
evitare di inizializzarla con ``make`` e si possono semplicemente
aggiungere elementi in coda uno alla volta, partendo dalla slice
vuota:

	var list []float64{} // slice vuota
	
	for {
		// ...
		list = append(list, elem)
	}

Questo approccio permette di gestire in maniera semplice un numero
variabile di input.

## Cicli

Per accedere agli elementi di una slice solitamente si possono usare i
due approcci che abbiamo visto con gli array, in generale si
preferisce il **for-range** perchè è più elegante.

	for i := 0; i < len(list); i++ {
		// len(list) restituisce la lunghezza di list
		fmt.Println(list[i])
	}
	
	for i := range list {
		fmt.Println(list[i])
	}

## Le stringhe

Per quelli di voi che si ricordano ciò che avevamo detto riguardo alle
stringhe, questo può sembrare molto simile: in effetti lo è poiché le
stringhe sono soltanto **slice di rune**.

# Esercizi
## Esercizio 0 - Database

*Problema*: Riprendete l'esercizio della volta scorsa e modificatelo
in modo che sia possibile specificare un numero arbitrario di voti per
ciascuna materia.

### Soluzione

{{cat database_slice.go | sed 's/^/	/g'}}

## Esercizio 1 - Fibonacci

*Problema*: Riprendete l'esercizio della volta scorsa e modificatelo
in modo che sia possibile specificare un intero ``n`` da terminale, ed
il programma generi una slice contenente i primi ``n`` numeri di
fibonacci.

### Soluzione

{{cat fib_slice.go | sed 's/^/	/g'}}
