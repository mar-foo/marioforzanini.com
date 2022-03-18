# Ripasso

Settimana scorsa abbiamo parlato di stringhe, caratteri, ASCII e
UTF-8, ``rune``...
Siccome il testo in generale è un argomento principe e necessario
della programmazione ho pensato di proporvi una serie di esercizi.

# Nota importante

L'ultima volta mi sono scordato di aggiungere che per leggere un carattere singolo da standard input non si può usare la funzione ``fmt.Scan``, occorre usare ``fmt.Scanf``. La sintassi è:

	var c byte
	fmt.Scanf("%c", &c)

oppure

	var r rune
	fmt.Scanf("%c", &r)


# Esercizio svolto: Cesare
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

## Soluzione

{{cat cesare1.go | sed 's/^/	/g'}}

## Soluzione 2

{{cat cesare2.go | sed 's/^/	/g'}}

# Esercizi

## Esercizio 0 - Conta maiuscole
*Problema*: Scrivere un programma go <a href="conta_maiuscole.go">conta_maiuscole.go</a> che legge
in input una stringa, conta il numero di maiuscole e stampa il
risultato.

*Esempi di esecuzione*:

Inserire una string: **AbCdEfG**

Maiuscole: 4

### Soluzione

{{cat conta_maiuscole.go | sed 's/^/	/g'}}

## Esercizio 1 - Capitalize
*Problema*: Scrivere un programma go ``capitalize.go`` che legge in
input una stringa e restituisce la stessa stringa in cui l'inizio di
ogni parola è maiuscolo mentre il resto è minuscolo.

*Esempi di esecuzione*:

Inserire una stringa: **nESSuno sCriVe iN quESTo moDo**

Nessuno Scrive In Questo Modo

### Soluzione

{{cat capitalize.go | sed 's/^/	/g'}}

## Esercizio 2 - Database
*Problema*: Scrivere un programma go <a href="database.go">database.go</a> che legge in
input una stringa che ha questo formato:

	nome;cognome;email;ultimo voto

e stampa i diversi campi su una riga differente con la legenda
corretta (è facoltativo aggiungere maiuscola al nome ed al cognome):

*Esempi di esecuzione*:

Inserire una stringa: **mario;forzanini;mario.forzanini@studenti.unimi.it;18**

Nome: Mario

Cognome: Forzanini

Email: mario.forzanini@studenti.unimi.it

Ultimo voto: 18

### Soluzione

{{cat database.go | sed 's/^/	/g'}}

## Esercizio 3 - Cripta password

*Problema*: Scrivere un programma <a href="cripta_password.go">cripta_password.go</a> che legge in
input una stringa e la 'trasforma' in una password applicando le
seguenti trasformazioni:

* 'a' -> @
* 'e' -> 3
* 'o' -> 0
* 's' -> $
* 'i' -> 1
* 'E' -> €
* 'L' -> £

*Esempi di esecuzione*:

Parola: **nomedelmiocane118**

Password: n0m3d3lm10c@n3118

*Nota*: Non usatelo per delle password vere

### Soluzione

{{cat cripta_password.go | sed 's/^/	/g'}}
