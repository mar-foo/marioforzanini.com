# Esercizi

Continuiamo con altri esercizi che
riguardano gli argomenti visti nelle
scorse lezioni: cicli **for** e
condizioni **if/else**.

## Esercizio 0 - Differenze

*Problema*: Scrivere un programma <a href="differenze.go">differenze.go</a>
che legge una serie di valori ``float64`` da tastiera e
stampa le differenze, cioè la differenza tra il secondo e il primo,
tra il terzo ed il secondo, e così via. Il programma termina quando
riceve in input il valore 0.


*Esempi di esecuzione*:

Scrivi un numero: **10**

Scrivi un numero: **20**

Differenza = 10

Scrivi un numero: **34**

Differenza = 14

Scrivi un numero: **0**

Differenza = -34

### Soluzione
{{cat differenze.go | sed 's/^/	/g'}}

## Esercizio 1 - Quadrati

*Problema*: Scrivere un programma <a href="quadrati.go">quadrati.go</a> che
legge in input un intero n positivo e calcola il massimo quadrato
``(k^2) <= n``.

*Esempi di esecuzione*:

Inserisci un numero: **56**

Massimo quadrato: 49

### Soluzione
{{cat quadrati.go | sed 's/^/	/g'}}

## Esercizio 2 - Euclide

*Problema*: Scrivere un programma <a href="euclide.go">euclide.go</a> che
legge da standard input due interi *a* e *b*, con ``a >=
b``, e calcola il MCD tra i due numeri con l'algoritmo di
Euclide.

*Esempi di esecuzione*:

Inserire due interi: **21 49**

MCD: 7

### Algoritmo di Euclide

Dati due numeri naturali *a* e *b*


0. Si controlla se *b** è zero
0. Se lo è, a è il MCD
0. Se non lo è, si assegna ad *r* il resto della divisione tra *a* e *b*, si pone ``a = b`` e ``b = r`` e
si ripete da 1.

## Esercizio 3 - Andamento

*Problema*: Scrivere un programma <a href="andamento.go">andamento.go</a>
che legge da tastiera una serie (di almeno un elemento) di numeri
interi > -1 e stampa "+" ogni volta che il nuovo valore è maggiore o
uguale al precedente e "-" altrimenti. Il programma si ferma quando il
numero in input è -1 e stampa la somma di tutti i numeri letti.

*Esempi di esecuzione*:

Inserire dei valori (terminare con -1): **2 4 7 3 9 5 -1**

++-+-+

somma: 31

## Esercizio 4 - Somma cifre

*Problema*: Scrivere un programma <a href="somma_cifre.go">somma_cifre.go</a>
che calcola la somma delle cifre di un numero intero positivo fornito
da standard input.

*Esempi di esecuzione*:

Inserire un numero: **5467913**

Somma: 35

## Esercizio 5 - Slash

*Problema*: Scrivere un programma <a href="slash.go">slash.go</a> che legge un intero positivo e stampa un
back-slash (\) di asterischi di altezza n.

*Esempi di esecuzione*:

Dimensione \: **3**

<pre>
*
 *
  *</pre>


### Soluzione
{{cat slash.go | sed 's/^/	/g'}}

## Esercizio 6 - Fibonacci

*Problema*: Scrivere un programma <a href="fibonacci.go">fibonacci.go</a>
che legge un intero positivo e stampa i numeri di fibonacci dal primo
all'n-esimo.

*Esempi di esecuzione*:

Inserire un numero: **6**

1 1 2 3 5 8

### Fibonacci

L'n-esimo numero della sequenza di fibonacci si ottiene sommando
l'(n-1)-esimo e l'(n-2)-esimo. I primi due numeri sono entrambe 1.
