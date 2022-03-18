# Correzione

Per prima cosa osserviamo le soluzioni agli
[esercizi precedenti](programmazione2.html)

# Ripasso

Nella scorsa lezione abbiamo parlato di lettura di dati e di
variabili, come due modi per rendere il programma dinamico ed
interattivo.

## Leggere input

Il costrutto per leggere input dall'utente ha questa forma:

	fmt.Scan(&nomeVariabile)

La variabile chiamata *nomeVariabile* deve essere stata dichiarata
precedentemente e l'utente deve inserire informazioni del tipo
corretto: se *nomeVariabile* è una variabile di tipo **int** non può
essere usata per conservare stringhe.

# Selezione ad una via: if

Oggi continuiamo con i concetti fondamentali della
programmazione e parliamo di una struttura che si ripeterà in molti
programmi: il costrutto if/else.

## Esempi

Molto spesso nei programmi occorre fare una scelta sulla base del
valore di qualche variabile o dell'avverarsi di qualche condizione.
Per esempio nell'[Esercizio 3](programmazione2.html) della volta scorsa sarebbe utile
decidere di abortire l'esecuzione se l'utente inserisce un voto che
non è compreso tra 3 e 10, oppure nell'[Esercizio 2](programmazione2.html) controllare se
la distanza totale e la quantità di carburante utilizzata sono
numeri diversi da zero, poichè dividere per zero fa andare in crash
il programma (provare per credere).

## Sintassi

La sintassi del costrutto **if/else** è composta da diverse parti:

* La parola chiave ``if``
* La condizione che va verificata
* Un blocco di codice (compreso tra '``{}``') da eseguire
nel caso in cui la condizione sia verificata

Per esempio per controllare che il contenuto di una variabile
chiamata *numero* sia uguale a zero si scriverebbe:


	if numero == 0 {
		// Cose da fare se uguale a 0
	}

In questo codice è apparso una nuova combinazione di simboli che
non avevamo ancora visto: l'operatore **==** che, come si può
intuire, ha questa sintassi: <primoArgomento> **==** <secondoArgomento>
e che si legge "primoArgomento uguale a secondoArgomento".

### Operatori di confronto

Gli operatori che confrontano diversi valori sono tanti, ma noi ne
vedremo 6. Ciascuno di questi argomenti restituisce un valore di
tipo **bool** (cioè *true* oppure *false*). Semplificando: ciascuno
di questi operatori risponde *si* oppure *no* ad una domanda sugli
argomenti che gli vengono forniti.

<table border="2" cellspacing="0" cellpadding="6" rules="groups" >
<thead>
<tr>
<th scope="col">Operatore</th>
<th scope="col">Domanda</th>
</tr>
</thead>
<tbody>
<tr>
<td>==</td>
<td>Il primo argomento è esattamente uguale al secondo?</td>
</tr>

<tr>
<td>!=</td>
<td>Il primo argomento è diverso dal secondo?</td>
</tr>

<tr>
<td><</td>
<td>Il primo argomento è minore del secondo?</td>
</tr>

<tr>
<td>></td>
<td>Il primo argomento è maggiore del secondo?</td>
</tr>

<tr>
<td><=</td>
<td>Il primo argomento è minore o uguale al secondo?</td>
</tr>

<tr>
<td>>=</td>
<td>Il primo argomento è maggiore o uguale al secondo?</td>
</tr>
</tbody>
</table>

#### Nota bene

Prestate attenzione a non confondere l'operatore di confronto ``==``
con l'operatore di assegnamento ``=``. Purtroppo l'abitudine di
utilizzare ``=`` per esprimere il fatto che due quantità siano
uguali può portare ad errori particolarmente ostici da trovare, se
volete confrontare due quantità ricordatevi di usare ``==``.

## Operatori booleani

Molto spesso occorre agire se più condizioni si verificano
contemporaneamente oppure se se ne verifica una sola, oppure se non
se ne verifica nessuna: ad esempio nell'[Esercizio 3](lezione2.html) della volta
scorsa vorremmo calcolare la media solo se primoVoto <= 10 **e**
primoVoto >= 3 Ma come faccio ad esprimere che voglio agire quando
sono vere l'una **e** l'altra cosa? Ho bisogno di nuovi simboli che
mi permettano di combinare dei valori booleani.

Occorre allora introdurre gli **operatori booleani**, ovvero quei
connettori logici che ci permettono di collegare diverse
condizioni. Questi 3 operatori sono facilmente traducibili in
italiano con delle espressioni di uso comune: **&&** (e), **||** (o),
**!** (non)


A questo punto nell'esercizio 3 della volta scorsa potrebbe
apparire un blocco di codice fatto così:


	if primoVoto <= 10 && primoVoto >= 3 && secondoVoto <= 10 && secondoVoto >= 3 {
		// Va tutto bene
	}

# Selezione a due vie: if/else

A volte occorre eseguire cose diverse se la condizione è verificata
oppure se non lo è, in questi casi al costrutto **if** si aggiungono
due parti:


* La parola chiave **else** (traduzione: *altrimenti*)
* Un blocco di codice da eseguire nel caso in cui la condizione non
sia verificata.

## Variabili locali

Una forma alternativa del costrutto **if** (oppure **if/else**)
consiste nel dichiarare una variabile prima di specificare la
condizione da verificare, ad esempio:



	if t := c * 2; t < 100 {
		fmt.Println("t vale", t, " che è minore di 100")
	} else {
		fmt.Println("t vale", t, " che è maggiore o uguale a 100")
	}

In questo caso la variabile t viene dichiarata subito dopo la
parola if e la condizione da testare segue il **;**.


Le variabili dichiarate in questo modo e le variabili dichiarate
all'interno dell'if sono locali, ovvero scompaiono non appena il
blocco di istruzioni è terminato.


	if t > 0 {
		s := t * 2 					// la variabile s è locale a questo blocco
		fmt.Println(s)
	}
	fmt.Println(s)					// ERRORE: La variabile s non esiste fuori dal blocco

# Selezione a più vie: if/else if

Se le condizioni mutualmente esclusive sono più di una si può usare
un costrutto simile ai precedenti:


	if <condizione> {
		// <condizione1> è vera
	} else if <condizione2> {
		// <condizione2> è vera
	} else {
		// nessuna delle precedenti è vera
	}

# Esercizi

## Esercizio 0 - Operatori di confronto e logici

*Problema*: Scrivere un programma go ``condizioni.go`` per testare,
una a una, le condizioni nella tabella che segue. Il programma, per
ogni condizione (a., &#x2026;, x.), legge un valore da tastiera (del
tipo indicato) e stampa ``true`` o ``false``, a seconda che la
condizione sia verificata o no. Implementare una condizione alla
volta, testarla su almeno due input (uno che la verifica ed uno che
la falsifica) e solo poi procedere alla successiva.

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">
<thead>
<tr>
<th scope="col">n</th>
<th scope="col">tipo</th>
<th scope="col">condizione</th>

</tr>
</thead>
<tbody>
<tr>
<td>a</td>
<td>int</td>
<td>uguale a 10</td>

</tr>

<tr>
<td>b</td>
<td>int</td>
<td>diverso da 10</td>
</tr>

<tr>
<td>c</td>
<td>int</td>
<td>diverso da 10 e da 20</td>
</tr>

<tr>
<td>d</td>
<td>int</td>
<td>diverso da 10 o da 20</td>
</tr>

<tr>
<td>e</td>
<td>int</td>
<td>maggiore o uguale a 10</td>
</tr>

<tr>
<td>f</td>
<td>int</td>
<td>compreso tra 10 e 20, inclusi</td>
</tr>

<tr>
<td>g</td>
<td>int</td>
<td>compreso tra 10 e 20, esclusi</td>
</tr>

<tr>
<td>h</td>
<td>int</td>
<td>compreso tra 10 e 20, 10 incluso e 20 escluso</td>
</tr>

<tr>
<td>i</td>
<td>int</td>
<td>minore di 10 o maggiore di 20</td>
</tr>

<tr>
<td>j</td>
<td>int</td>
<td>tra 10 e 20 inclusi o tra 30 e 40 inclusi</td>
</tr>

<tr>
<td>k</td>
<td>int</td>
<td>multiplo di 4 ma non di 100</td>
</tr>

<tr>
<td>l</td>
<td>int</td>
<td>dispari e compreso tra 0 e 100, inclusi</td>
</tr>
</tbody>
</table>

*Esempi di esecuzione:*

int uguale a 10: **10**

true

int uguale a 10: **11**

false

*Nota*: Il seguente codice:


	if n == 10 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

si può scrivere in modo più compatto (ed equivalente) così:


	fmt.Println(n == 10)


Questo perchè l'espressione **n == 10** quando viene valutata dà come
risultato ``true`` se è vera e ``false`` altrimenti.

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var n int
		
		// Punto a
		fmt.Println("int uguale a 10: ")
		fmt.Scan(&n)
		fmt.Println(n == 10)
		
		// Punto b
		fmt.Println("int diverso da 10: ")
		fmt.Scan(&n)
		fmt.Println(n != 10)
		
		// Punto c
		fmt.Println("int diverso da 10 e da 20: ")
		fmt.Scan(&n)
		fmt.Println(n != 10 && n != 20)
		
		// Punto d
		fmt.Println("int diverso da 10 o da 20: ")
		fmt.Scan(&n)
		fmt.Println(n != 10 || n != 20)
		
		// Punto e
		fmt.Println("int maggiore o uguale a 10: ")
		fmt.Scan(&n)
		fmt.Println(n >= 10)
		
		// Punto f
		fmt.Println("int compreso tra 10 e 20, inclusi: ")
		fmt.Scan(&n)
		fmt.Println(n <= 20 && n >= 10)
		
		// Punto g
		fmt.Println("int compreso tra 10 e 20, esclusi: ")
		fmt.Scan(&n)
		fmt.Println(n < 20 && n > 10)
		
		// Punto h
		fmt.Println("int compreso tra 10 e 20, 10 incluso e 20 escluso: ")
		fmt.Scan(&n)
		fmt.Println(n < 20 && n >= 10)
		
		// Punto i
		fmt.Println("int minore di 10 o maggiore di 20")
		fmt.Scan(&n)
		fmt.Println(n < 10 || n > 20)
		
		// Punto j
		fmt.Println("int compreso tra 10 e 20 o tra 30 e 40, estremi inclusi: ")
		fmt.Scan(&n)
		fmt.Println((n <= 20 && n >= 10) || (n <= 40 && n >= 30))
		
		// Punto k
		fmt.Println("int multiplo di 4 ma non di 100")
		fmt.Scan(&n)
		fmt.Println(n % 4 == 0 && n % 100 != 0)
		
		// Punto l
		fmt.Println("int dispari e compreso tra 0 e 100, inclusi")
		fmt.Scan(&n)
		fmt.Println(n % 2 == 1 && 0 <= n && a <= 100)
		
		return
	}

## Esercizio 1 - Voto valido

*Problema:* Scrivere un programma Go ``voto_valido.go`` che legge un
numero intero, se il numero non è compreso tra 3 e 10, stampa "voto
non valido", altrimenti non stampa niente.


*Esempi di esecuzione*:

voto: **11**

voto non valido

voto: **4**

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var voto int
		fmt.Println("voto: ")
		fmt.Scan(&voto)
		
		if voto < 3 || voto  > 10 {
			fmt.Println("voto non valido")
		}
		
		return
	}

## Esercizio 2 - Maggiore

*Problema:* Scrivere un programma Go maggiore.go che legga due
interi, li salvi in due variabili ``max`` e ``min`` in qualsiasi
ordine; se non sono in ordine, li sistemi in modo che ``min``
contenga il minore e ``max`` il maggiore; infine stampi il contenuto
di ``max``.


*Esempi di esecuzione*:

due int: **10 20**

20

due int: **20 10**

20

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var max, min int
		
		fmt.Println("due int:")
		fmt.Scan(&max)
		fmt.Scan(&min)
		
		if min > max {
			max, min = min, max // Scambio i valori usando l'assegnamento multiplo
		}
		
		fmt.Println(max, min)
		
		return
	}

## Esercizio 3 - Pari e dispari

*Problema:* Scrivere un programma ``pari_dispari.go`` che legge un
intero ``n`` e, a seconda del valore di ``n``, stampa uno dei messaggi
"n è pari" oppure "n è dispari".


*Esempi di esecuzione*:

numero: **4**

4 è pari

numero: **5**

5 è dispari


*Nota*: Un numero intero ``n`` è pari se il resto della divisione 
di ``n`` per 2 è pari a zero. In go l'operatore che calcola il resto della
divisione intera è ``%``.

### Soluzione

	package main
	
	import "fmt"
	
		func main() {
		var numero int
		
		fmt.Println("numero:")
		fmt.Scan(&numero)
		
		if numero % 2 == 0 {
			fmt.Println(numero, " è pari")
		} else {
			fmt.Println(numero, " è dispari")
		}
		
		return
	}

## Esercizio 4 - Tariffe scontate

*Problema:* Scrivere un programma ``tariffe.go`` che chiede
all'utente l'età (``int``) e se è studente (``bool``) e stampa il costo
del biglietto di ingresso al cinema secondo la seguente tabella:


<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">
<thead>
<tr>
<th scope="col">età</th>
<th scope="col">tariffa</th>
</tr>
</thead>
<tbody>
<tr>
<td>[0-9) anni</td>
<td>gratis</td>
</tr>

<tr>
<td>[9-14) anni</td>
<td>5</td>
</tr>

<tr>
<td>[14-26) anni</td>
<td>7.5</td>
</tr>

<tr>
<td>[26-65) anni</td>
<td>10</td>
</tr>

<tr>
<td>>=65 anni</td>
<td>7.5</td>
</tr>

<tr>
<td>studenti >= 14</td>
<td>5</td>
</tr>
</tbody>
</table>


*Esempi di esecuzione*:

età: 16

studente? (t/f): t

ingresso 5 euro

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var eta int
		var studente bool
		
		fmt.Println("età:")
		fmt.Scan(&eta) 
		
		fmt.Println("studente? (t/f):")
		fmt.Scan(&studente)
		
		if eta >= 0 && eta < 9 {
			fmt.Println("gratis")
		} else if studente || eta < 14 {
			fmt.Println("5 euro")
		} else if (eta >= 14 && eta < 26) || eta > 65 {
			fmt.Println("7.5 euro")
		} else {
			fmt.Println("10 euro")
		}
		
		return
	}

## Esercizio 5 - Sovrapposizione

*Problema:* Scrivere un programma ``sovrapposizione.go`` che legge da
tastiera il giorno [1-31], l'ora di inizio [0-24] e l'ora di fine
[0-24] di due appuntamenti e stabilisce se si sovrappongono (anche
parzialmente) oppure no.


*Esempi di esecuzione*:

appuntamento 1 (gg, start, end): **28 15 18**

appuntamento 2 (gg, start, end): **28 16 20**

si sovrappongono

appuntamento 1 (gg, start, end): **11 15 18**

appuntamento 2 (gg, start, end): **28 16 20**

non si sovrappongono

### Soluzione

	package main
	
	import "fmt"
	
	func main() {
		var gg1, start1, end1 int
		var gg2, start2, end2 int
		
		fmt.Println("appuntamento 1 (gg, start, end): ")
		fmt.Scan(&gg1, &start1, &end1)
		
		fmt.Println("appuntamento 2 (gg, start, end): ")
		fmt.Scan(&gg2, &start2, &end2)
		
		if gg1 != gg2 {
			fmt.Println("non si sovrappongono")
		} else {
			if !(start2 < end1 && start1 < end2) {
				fmt.Println("non si sovrappongono")
			} else {
				fmt.Println("si sovrappongono")
			}
		}
		return
	}
