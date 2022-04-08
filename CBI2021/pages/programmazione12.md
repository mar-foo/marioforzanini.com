# Ripasso

Settimana scorsa abbiamo parlato di gestione dei file. Fortunatamente,
la gestione dei file è molto semplice e bisogna imparare soltanto 3
funzioni:

	os.Open(name string) (*os.File, error)
	os.Create(name string) (*os.File, error)
	os.OpenFile(name string, flag int, perm FileMode) (*os.File, error)

e ricordarsi di chiudere il file una volta aperto.

# Gestione degli errori

Oggi impariamo qualcosa che sarà molto utile in tutti i programmi che
scriveremo da qui in avanti, e che ritornerà molto spesso anche
nell'utilizzo della libreria standard di Go, o di librerie esterne. Si
tratta della gestione degli errori.

Ogni funzione, come abbiamo visto, può avere più di un valore di
ritorno e molto spesso le funzioni che abbiamo usati ritornano due
valori, per esempio:

	os.Open(name string) (*os.File, error)

ritorna due valori: uno di tipo ``*os.File`` e l'altro di tipo
``error``. Come si può intuire il secondo valore sarà ciò su cui
concentreremo la nostra attenzione oggi.

## Error

È uso comune tra i programmatori go utilizzare una particolare
struttura per gestire gli errori: ogni volta che la funzione che si
sta sviluppando potrebbe incontrare un errore, questo viene catturato
in una variabile di tipo ``error`` che viene poi restituita alla fine
della funzione per essere gestita. Le variabili di tipo ``error``
possono essere stampate, il loro contenuto sarà il messaggio di
errore, per esempio se cercate di aprire un file inesistente:

	f, err := os.Open("inesistente.txt") // err conterrà l'errore
	// se non c'è nessun errore il valore di err sarà nil
	if err != nil { // significa che c'è stato qualche errore
		fmt.Fprintln(os.Stderr, err) // usiamo lo standard error!
		return // blocchiamo l'esecuzione del programma
	}

Il risultato di questo programma sarà:

	open inesistente.txt: no such file or directory

### Creare un error

Per utilizzare questo idioma nelle nostre funzioni dovremo
usare la funzione

	Errorf(format string, a ...interface{}) error

che funziona come ``fmt.Printf``, cioè accetta una stringa che
specifica come formattare l'errore, ma invece di stampare il risultato
a video lo ritorna con tipo errore. Ad esempio una funzione che non fa
nulla ma ritorna un errore:

	func alwaysError() error {
		var err error
		err = fmt.Errorf("So solo sbagliare\n")
		return err
	}

oppure

	func alwaysError() error {
		return fmt.Errorf("So solo sbagliare\n")
	}

## Gestire un error

Quando vi ritrovate a lavorare con funzioni che ritornano degli errori
(vedi funzioni della famiglia di ``fmt.Scan``), potete decidere cosa
fare nel momento in cui incontrate un errore.

Solitamente si distingue tra due tipi di errori: errori dell'utente ed
errori del programmatore. Una buona regola è cercare di risolvere gli
errori dell'utente e bloccare l'esecuzione del programma nel caso di
errori del programmatore. Per esempio, se l'utente inserisce il nome
di un file che non esiste si può decidere di ignorare quel file e
proseguire notificando l'utente.

	var filename string
	fmt.Print("File: ")
	fmt.Scan(&filename)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Il file %s non esiste!\n", filename)
	}
	// proseguo comunque

## Tipi di errore

``error`` permette anche di specificare il *tipo* di errore che si
vuole ritornare. Per esempio, il pacchetto ``io`` definisce il tipo di
errore ``io.EOF`` che è quello che ritornano le funzioni tipo
``fmt.Scan`` quando raggiungono la fine di un file, questo era utile
per gli esercizi della scorsa volta:

	err := fmt.Fscanf("%c", &c)
	if err == io.EOF {
		// Il file è finito
	}

## Funzioni che ritornano bool

Un altro modo di gestire gli errori è quello di restituire un valore
booleano al posto di un ``error``. Solitamente questo valore booleano
viene interpretato come un **ok**, ovvero: se il valore è ``true``
significa che non ci sono stati errori, altrimenti qualcosa è andato
storto. Per esempio:

	func main() {
		s, ok := neverOk()
		if ! ok {
			fmt.Fprintf("Qualcosa è andato storto: %s\n", s)
		}
		return
	}

	func neverOk() (string, bool) {
		return "Non è tutto ok", false
	}

Questo tipo di gestione degli errori sta progressivamente
sparendo vista la maggiore flessibilità del tipo ``error`` (che per
esempio permette di descrivere quale sia stato l'errore e di
specificare diversi tipi di errore).

# Array

Dato che la gestione degli errori è un argomento interessante ma poco
applicativo, occorre aggiungere qualcosa per poter fare esercizi oggi.
Iniziamo quindi a parlare di **collezioni**, ovvero strutture che
contengono un certo numero di elementi. Molti di questi sono comuni
ad altri linguaggi di programmazione.

Il primo tipo di struttura che trattiamo sono gli array: contenitori
di un numero fissato di elementi dello stesso tipo. Gli array sono
indicati con la notazione ``[n]`` dove ``n`` è il numero di elementi.
Per esempio, l'esercizio della media di 10 numeri potrebbe essere
risolto con un array di 10 ``float64``:

	var voti [10]float64

la variabile voti rappresenta un contenitore per 10 ``float64``.
Questo tipo di collezione è poco flessibile:

* Il numero di elementi è fissato, in particolare va conosciuto prima di compilare
* ``[5]float64`` è un tipo diverso da ``[10]float64``

## Accesso agli elementi

Per accedere alle variabili che compongono un array si usa
l'operazione di **indicizzazione**: ogni elemento dell'array occupa
una posizione che varia tra 0 e la lunghezza - 1. La sintassi è la
seguente:

	var voti [10]float64
	voti[0] = 5 // Il primo voto sarà 5

Se si cerca di accedere ad un elemento che non è nel range corretto
(``>= len(array)``) il programma va in crash (il compilatore non
se ne accorge).

## Listare elementi di un array

Per accedere a tutti gli elementi di un array si usano cicli ``for``,
di due tipi: semplici cicli ``for``

	var array1 [5]int
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

e cicli ``for range``:

	var array2 [55]int
	
	for i := range array2 { // i è un intero contenente l'indice
		fmt.Println(array2[i])
	}

questi costrutti oltre che listare gli elementi permettono anche di
cambiarli: gli array sono strutture read-write.

	var array3 [3]int
	
	for i := 0; i < len(array3); i++ {
		array3[i] = i + 1
	}

``array3`` conterrà i numeri 1, 2, 3.

## Passare gli array come argomenti di funzioni

Come detto sopra, la lunghezza dell'array ne cambia il tipo, quindi
nella dichiarazione di una funzione che ha come input un array dovete
specificarla. Per esempio la funzione che calcola la media di 10 voti
contenuti in un array può essere dichiarata così:

	func media(voti [10]float64) float64

# Esercizi

## Esercizio 0 - Media di 10 voti

*Problema*: Implementare il solito problema della media dei voti
utilizzando un array di tipo ``[10]float64``. Leggete i voti da un
file **voti.txt** che contenga 10 numeri.

### Soluzione

{{cat media_array.go | sed 's/^/	/g'}}

## Esercizio 1 - Database 3

*Problema*: Scrivere un programma ``database.go`` che legge da un file
**database.txt** che abbia questa sintassi:

	materia voto

e calcola la media dei voti associati ad una particolare materia. Le
materie possibili sono 3: **matematica**, **fisica** e
**programmazione**, a ciascuna di queste associate un array di 10
``float64`` che rappresentano i voti. Stampate le medie associate a
ciascuna delle 3 materie. Ricordatevi di controllare se i voti sono
nel range corretto e gestite l'errore se non lo sono. Spezzate il
codice in funzioni per aiutarvi a mantenerlo ordinato.

*Nota*: Penso che valga la pena provare a fare bene almeno questo
esercizio.

*Esempio di esecuzione*:

	medie:
	matematica 5.5
	fisica 6
	programmazione 10

il file database.txt:

	matematica 4
	matematica 7
	fisica 3
	programmazione 10
	fisica 6
	matematica 5.5
	... 10 voti per ogni entrata

### Soluzione

{{cat database_3.go | sed 's/^/	/g'}}

## Esercizio 2 - Fibonacci

*Problema*: Scrivere un programma go ``fib.go`` che calcola i primi 10
numeri della sequenza di fibonacci. Utilizzate un array di 10 interi e
ricordatevi che vi servono sempre 2 interi della sequenza per
calcolare quello successivo.

*Definizione*: I primi due numeri della sequenza di fibonacci sono 0 e
1, tutti gli altri numeri sono la somma dei due precedenti nella
sequenza.

### Soluzione

{{cat fibonacci.go | sed 's/^/	/g'}}
