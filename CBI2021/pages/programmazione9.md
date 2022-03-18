# Byte e Rune

Go mette a disposizione due tipi per descrivere un carattere: ``byte``
e ``rune``.  Entrambe questi tipi vengono rappresentati in memoria
come degli interi.

## Byte

Il tipo ``byte`` viene rappresentato in memoria come un intero di 8
bit senza segno (``uint8``), i cui valori sono associati univocamente
a lettere secondo la tabella ASCII.  I letterali ``byte`` sono
circondati da ``'``.

### La codifica ASCII

Lo standard ASCII risale al 1961 e contiene le indicazioni su come
codificare lettere e simboli dell'alfabeto americano come interi.
Ogni carattere occupa un byte.  Poiché in inglese non sono presenti
accenti nè altri simboli utilizzati in altre lingue, questa codifica è
stata usata come base per un altro standard più versatile: UTF-8.

## Rune

Il tipo ``rune`` rappresenta un ``code point`` della codifica UTF-8,
questo tipo è quello che useremo per la codifica dei caratteri nei
nostri sorgenti.

### UTF-8 (Unicode)

La codifica UTF-8 utilizza una quantità variabile di bit per
codificare un range di 1.112.064 ``code points``.  I primi 128
caratteri sono quelli della tabella ASCII, gli altri caratteri
spaziano da lettere accentate fino ad emoticon.  Per quanto capace di
rappresentare caratteri che necessitano di 4 byte (32 bit), i
caratteri ASCII utilizzano comunque 1 byte (8 bit).

### Il pacchetto unicode

Il pacchetto ``unicode`` contiene molte funzioni utili per lavorare
con ``rune``, per esempio:

	func IsLetter(r rune) bool
	func IsDigit(r rune) bool
	func IsSpace(r rune) bool
	func ToLower(r rune) rune

# Stringhe

In Go una stringa è una sequenza di ``rune``, cioè una sequenza di
caratteri di grandezza variabile.  Le stringhe sono un tipo
immutabile, una volta create non possono essere modificate (capiremo
più avanti perchè).

Ci sono 2 tipi di letterali stringa:

* **Interpretato**: Circondato da ``""``, le sequenze di escape vengono interpretate (per esempio ``\n`` rappresenta un 'a capo', ``\t`` rappresenta un tab, ``\"`` rappresenta letteralmente un ")
* **Raw**: Circondato da \`\`, vengono prese letteralmente

Per esempio:

	str := "Questa è la prima linea\nQuesta è la seconda linea"
	str2 := ``Questa è la prima linea
	Questa è la seconda``

Il valore di default di una stringa è ``""``.

## La funzione len()

La lunghezza di una stringa ``str`` si ottiene con la funzione
``len``:
	len(str)

## Concatenare stringhe

Due stringhe ``s1`` ed ``s2`` possono venire concatenate in un'unica
stringa utilizzando ``+``

	s := s1 + s2

## For-range

Go mette a disposizione un quarto tipo di loop che agisce sulle
stringhe di ``rune``.  Vista la natura variabile in spazio dei
caratteri UTF-8 non è possibile trattarli con un semplice ``for``
ternario.  La sintassi è la seguente:

	for pos, byte := range str {
		// pos contiene l'indice attuale, byte la runa attuale
	}

Con questo costrutto è possibile iterare sui ``code points`` uno alla
volta, ``byte`` conterrà una copia della runa attuale (è utile solo
per leggere valori, non può essere usato per cambiare la stringa).

## Strconv e strings

I pacchetti ``strconv`` e ``strings`` contengono molte funzioni utili per lavorare sulle stringhe.

### Esempi

#### strconv.Atoi

	func Atoi(s string) (int, error)

Ritorna come primo valore l'intero corrispondente alla stringa s:

	package main

	import (
		"fmt"
		"strconv"
	)

	func main() {
		var str string
		var n, m int

		n = 50
		str = "45"

		m, _ = strconv.Atoi(str)
		fmt.Println(m + n) // Stampa 95
		return
	}

La funzione ``Itoa`` fa l'opposto:

	func Itoa(i int) string

#### strings.HasPrefix, strings.HasSuffix

	func HasPrefix(s, prefix string) bool
	func HasSuffix(s, suffix string) bool

Queste funzioni testano se la stringa ``s`` contiene come
suffisso/prefisso la stringa ``prefix`` o ``suffix``.  Siccome
ritornano un ``bool`` possono essere usate in if e for.

#### strings.Contains

	func Contains(s, substr string) bool

La funzione ``strings.Contains`` controlla se la stringa ``s``
contiene la sottostringa ``substr``.

#### E molte altre...

# Esercizi

Nella risoluzione di questi esercizi mi aspetto 2 cose:
0. Utilizzate ``go doc`` per trovare funzioni che potrebbero servirvi nei pacchetti ``unicode``, ``strings`` e ``strconv``
0. Cercate di utilizzare delle funzioni per rendere il codice dentro ``main()`` più leggibile. Ad esempio nell'esercizio 2 potreste servirvi di una funzione ``func isVocale(r rune) bool`` e di una funzione ``func contaVocali(s string) int``

## Esercizio 0
*Problema*: Scrivere un programma <a href="es0.go">es0.go</a> che legge un byte (Occorre
uno Scan in più per catturare l'invio) e
0. Stampa il byte precedente, il byte stesso e quello successivo in ordine lessicografico (ASCII)
0. Stabilisce se è una lettera tra A e L (maiuscole), o altro

*Nota*: Siccome i ``byte`` vengono rappresentati come interi in
memoria, ``b`` può anche essere scritto come

	'a' + 1

*Esempio di esecuzione*:

Inserire un carattere: **n**

mno

altro

Inserire un carattere: **C**

BCD

A-L

### Soluzione

{{cat es0.go | sed 's/^/	/g'}}

## Esercizio 1 - Rompi stringa
*Problema*: Scrivere un programma <a href="rompi_stringa.go">rompi_stringa.go</a> che legge in
input una stringa e la stampa in verticale, una runa per riga.

*Esempio di esecuzione*:

Inserire una stringa: **Perchè**

P
<br>
e
<br>
r
<br>
c
<br>
h
<br>
è

### Soluzione

{{cat rompi_stringa.go | sed 's/^/	/g'}}

## Esercizio 2 - Conta vocali

*Problema*: Scrivere un programma <a href="conta_vocali.go">conta_vocali.go</a> che legge in
input una parola e stampa il numero di lettere (``rune``) che sono
vocali (aeiou).

*Esempio di esecuzione*:

Inserire una parola: **Golang**
<br>
2

Inserire una parola: **supercalifragilisticoespiralidoso**

15

### Soluzione

{{cat conta_vocali.go | sed 's/^/	/g'}}

## Esercizio 3 - Trova rune

*Problema*: Scrivere un programma <a href="trova.go">trova.go</a> che legge un ``rune`` e
una stringa e stampa la prima posizione (indicizzata a partire da 0)
del carattere nella stringa, o -1 se il carattere non c'è.

*Esempio di esecuzione*:

Carattere da cercare: **c**

Stringa: **Tre tigri contro tre tigri**

10

### Soluzione

{{cat trova.go | sed 's/^/	/g'}}

## Esercizio 4 - Minuscole o maiuscole

*Problema*: Scrivere un programma <a href="minuscole_maiuscole.go">minuscole_maiuscole.go</a> che legge
una stringa e stabilisce se questa contiene solo maiuscole, solo
minuscole o entrambe e stampa un messaggio opportuno.

*Nota*: Il pacchetto ``unicode`` contiene le funzioni ``IsLower`` e
``IsUpper`` che potrebbero essere utili

*Esempio di esecuzione*:

Inserire una stringa: **ATTENZIONE**

Solo maiuscole

Inserire una stringa: **Una frase**

Sia maiuscole che minuscole

Inserire una stringa: **dopo un punto fermo occorre una maiuscola**

Solo minuscole

### Soluzione

{{cat minuscole_maiuscole.go | sed 's/^/	/g'}}

## Esercizio 5 - Cesare
*Problema*: Scrivere un programma <a href="cesare.go">cesare.go</a> che legge un valore
``int`` non negativo ``k`` (la chiave di cifratura) e una sequenza di
lettere minuscole (sulla stessa riga e senza spazi), terminate da
'\n'.  Il programma stampa la sequenza letta cifrata secondo il
cifrario di Cesare, usando come chiave ``k``: ogni lettera del testo
viene sostituita dalla lettera che si trova ``k`` posizioni dopo
nell'alfabeto, ritornando dopo la 'z' alla lettera 'a'.

*Nota*: Ricordate che i ``byte`` sono semplicemente numeri.  Quindi
per rimanere entro un certo range si può usare l'operazione '%': il
resto della divisione intera per ``n`` sta tra ``0`` e ``n - 1`` e la
lunghezza dell'alfabeto è ``'z' - 'a' + 1``.

*Esempio di esecuzione*:

chiave: **2**

testo: **zaprb**

bcrtd

chiave: **100**

testo: **abcd**

wxyz

### Soluzione

Vedi prossima lezione
