# Correzione

Per prima cosa osserviamo le soluzioni agli [esercizi precedenti](programmazione1.html).

# Ripasso
Nella scorsa lezione ciò che abbiamo capito del linguaggio è che per
stampare a video un messaggio occorre utilizzare una costruzione
fatta così:

	package main
	import "fmt"

	func main() {
		fmt.Println(<messaggio>)
		return
	}

In realtà ciò che si occupa di stampare il messaggio è soltanto la riga

	fmt.Println(<messaggio>)

tutto il resto è una construzione che si ripeterà sempre uguale nei
nostri programmi e che serve a "far partire" il programma nel modo
giusto.

## Le variabili

In go ciascun dato che vogliamo utilizzare va racchiuso in un
*contenitore*, detto **variabile**. Ogni variabile ha un **tipo**
(ovviamente se voglio rappresentare un numero ho bisogno di un
contenitore di un certo tipo, se voglio rappresentare una parola di
un altro). Inizialmente ci occuperemo soltanto di numeri perchè i
contenitori per le stringhe sono un pochino più complicati.

### Dichiarare variabili

La parola chiave di Go per dichiarare una variabile è **var**,
la sintassi è questa:

	var nome tipo

Per dichiarare una variabile occorre darle un nome, in
questo caso **n**; ogni volta che nel programma ci si
riferisce ad **n** si sta parlando del valore che sta nel
'contenitore' di nome **n**. Per mettere un valore dentro
**n** si usa l'operatore ``=``. Ad esempio il
seguente programma dichiara una variabile di tipo intero ``n``,
mette 10 dentro ``n``e poi stampa il contenuto della
variabile.

	package main
	import "fmt"

	func main() {
var n int
n = 10
fmt.Println(n)
return
	}

### Tipi di variabili

Esistono diversi tipi di variabili definiti da Go:
``int``, ``int64``, ``uint``,...
Noi inizialmente ci occuperemo principalmente di 3 tipi:
``int``, ``float64``, ``bool``; utili
per rappresentare, rispettivamente, numeri interi, numeri
decimali e valori di verità (*true* oppure *false*).

	package main
	import "fmt"

	func main() {
var n float64
n = 1.10
fmt.Println(n)
return
	}

## Nota: dichiarazioni multiple

In una sola riga posso dichiarare più variabili dello stesso tipo,
la sintassi è:

	var nome1, nome2 tipo

## Nota: omogeneità

In Go ogni espressione può contenere solo variabili dello stesso
tipo: per esempio non posso sommare un intero ed un numero
decimale. Lo stesso programma di prima, ma dichiarando le due
variabili come intero e come float64 dà un errore.

	package main
	import "fmt"

	func main() {
var n int
var m float64
n = 10
m = 20.0
fmt.Println(n + m) // ERRORE
return
	}

## Nota: variabili inizializzate

Molto spesso risulta utile dare valori iniziali alle variabili, per
farlo occorre utilizzare l'operatore '``:=``' tra il nome della variabile
ed il valore con cui la si vuole inizializzare. Go cercherà di capire
il tipo della variabile e assegnerà il valore.

	package main
	import "fmt"

	func main() {
primo := 20
fmt.Println(primo)
return
	}

## Nota: identificatore blank

A volte occorre ignorare alcuni valori, poichè non vale la pena di
salvarli in una variabile oppure perchè non verranno mai usati nel
resto del programma. Per fare questa cosa, al posto del nome della
variabile occorre mettere il simbolo ``_``

	package main
	import "fmt"

	func main() {
_, n := 10, 20
fmt.Println(n)
return
	}

# Lettura da tastiera

L'argomento di oggi è la lettura di dati dall'utente, tutte le
interazioni avvengono attraverso il terminale, quindi anche la
lettura di valori avviene quando l'utente scrive qualcosa nel
terminale.
Leggere valori da tastiera nello specifico significa salvare
l'input dell'utente in una variabile che abbia il tipo adatto. Se
ci aspettiamo che l'utente inserisca un numero intero, prepariamo
una variabile di tipo ``int``, se ci aspettiamo che l'utente
inserisca un numero con la virgola una variabile di tipo ``float64``.
Vedremo successivamente come controllare che l'utente abbia
inserito un valore del tipo corretto, per ora diamo per scontato
che sia così.

Per leggere valori dall'utente occorre utilizzare un costrutto
molto simile a quello che utilizzavamo per scrivere in output, ad
esempio un programma che legge da tastiera un numero intero e
conserva il valore letto in una variabile è fatto così:

	package main
	import "fmt"

	func main() {
var numero int
fmt.Scan(&numero)
return
	}

Alla fine del programma nella variabile ``numero`` sarà conservato il
valore inserito dall'utente. Per inserire il valore letto
all'interno di una variabile occorre quindi utilizzare un costrutto
che ha la forma:

	fmt.Scan(&<nome variabile>)

## Nota utile per gli esercizi

L'operazione di divisione è definita in modo diverso per interi e
per numeri decimali: su valori di tipo ``int`` l'operazione è quella
della divisione intera: il risultato è un numero intero,
eventualmente si può calcolare il resto con l'operazione ``%`` (detta
**modulo**); su valori di tipo ``float64`` l'operazione è quella usuale.

# Esercizi

## Esercizio 0 - Conversione da secondi in minuti

*Problema*: Scrivere un programma Go <a href="secondi_minuti.go">secondi_minuti.go</a> che, preso
in input un valore **intero** rappresentante una quantità in secondi,
stampi in output il corrispondente valore in minuti.

*Esempio di esecuzione*:

Numero di secondi: **120**

Numero di minuti: **2**

### Soluzione
{{cat secondi_minuti.go | sed 's/^/	/g'}}

## Esercizio 1 - Conversione da ore in giorni


*Problema:* Scrivere un programma Go <a href="minuti_ore.go">minuti_ore.go</a> che, preso in
input un valore **intero** rappresentante una quantità in ore,
stampi in output il corrispondente valore in giorni.


*Esempio di esecuzione*

Numero di ore: **48**

Numero di giorni: 2

### Soluzione
{{cat minuti_ore.go | sed 's/^/	/g'}}

## Esercizio 2 - Consumo medio e resa di un motore

*Problema:* Scrivere un programma Go <a href="consumo_resa.go">consumo_resa.go</a> che calcola
il consumo medio e la resa di un motore date la distanza totale
percorsa (in km) e la quantità di carburante utilizzata (in litri).
I valori in input sono di tipo float64.

*Nota:* Il consumo medio di carburante si esprime in l/km ed è la
quantità di carburante che occorre in media per percorrere un km di
strada. La resa di un motore è data dalla distanza percorsa in
media con un litro di carburante e si esprime con km/l.

*Esempio di esecuzione:*

Distanza percorsa (in km): **50**

Quantità di carburante utilizzata (in l): **10**

Consumo medio: 0.2 l/km

Resa media: 5 km/l

### Soluzione
{{cat consumo_resa.go | sed 's/^/	/g'}}

## Esercizio 3 - Media di due voti

*Problema:* Scrivere un programma Go <a href="media_due_voti.go">media_due_voti.go</a>
che, letti due voti (di tipo float64), ne calcola la media.


*Esempio di esecuzione:*

Primo voto: **6**

Secondo voto: **7**

Media: 6.5
### Soluzione
{{cat media_due_voti.go | sed 's/^/	/g'}}

## Esercizio 4 - Velocità

*Problema:* Scrivere un programma Go <a href="velocita.go">velocita.go</a> che calcola la
velocità media su un tratto di strada, dopo aver letto in input la
lunghezza del tratto (in km) e il tempo trascorso nel percorrerlo
(in ore). I valori letto sono di tipo float64.



*Nota:* La velocità media si esprime in km/h e si calcola dividendo
la lunghezza dell tratto percorso per il tempo necessario a
percorrerlo


*Bonus:* Utilizzando i risultati di conversione degli esercizi
precedenti, come ricavereste la velocità in m/s.



*Esempio di esecuzione:*

Lunghezza del percorso (in km): **72**

Tempo trascorso (in h): **2**

Velocità media: 36 km/h

Velocità media: 10 m/s

### Soluzione
{{cat velocita.go | sed 's/^/	/g'}}

## Esercizio 5 - Conversione da secondi a giorni, ore, minuti

*Problema:* Scrivere un programma Go <a href="secondi_giorni.go">secondi_giorni.go</a>
che, preso in input un valore *intero* rappresentante una
quantità in secondi, stampi in output il corrispondente valore
in giorni, ore, minuti, secondi.



*Bonus:* Provate a farlo senza usare mai l'operazione di sottrazione.



*Esempio di esecuzione:*

Numero di secondi: **123456**

g:h:m:s = 1:10:17:36

### Soluzione
{{cat secondi_giorni.go | sed 's/^/	/g'}}
