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
*Problema*: Scrivere un programma ``cesare.go`` che legge un valore
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

	package main

	import (
		"fmt"
		"strings"
	)

	func main() {
		var key int
		var s string

		fmt.Print("Chiave: ")
		fmt.Scan(&key)
		fmt.Print("testo: ")
		fmt.Scan(&s)

		fmt.Println(caesar(s, key))
		return
	}

	func caesar(s string, key int) string {
		var result string
		s = strings.ToLower(s)
		for _, c := range s {
			result += string('a' + (int(c) + key - 'a') % ('z' - 'a' + 1))
		}
		return result
	}

## Soluzione 2

	package main

	import "fmt"

	func main() {
		var key int
		var c byte

		fmt.Print("Chiave: ")
		fmt.Scan(&key)
		fmt.Print("testo: ")
		for fmt.Scanf("%c", &c); c != '\n'; fmt.Scanf("%c", &c) {
			fmt.Print(string(caesar(c, key)))
		}
		fmt.Println()
		return
	}

	func caesar(c byte, key int) byte {
		return byte('a' + (int(c) + key - 'a') % ('z' - 'a' + 1))
	}

# Esercizi

## Esercizio 0 - Conta maiuscole
*Problema*: Scrivere un programma go ``conta_maiuscole.go`` che legge
in input una stringa, conta il numero di maiuscole e stampa il
risultato.

*Esempi di esecuzione*:

Inserire una string: **AbCdEfG**

Maiuscole: 4

### Soluzione

	package main

	import (
		"fmt"
		"strings"
	)

	func main() {
		var input string
		fmt.Print("Inserire una stringa: ")
		fmt.Scan(&input)
		fmt.Println("Maiuscole: ", contaMaiuscole(input))
	}

	func contaMaiuscole(s string) int {
		var count int
		count += strings.Count(s, "a")
		count += strings.Count(s, "e")
		count += strings.Count(s, "i")
		count += strings.Count(s, "o")
		count += strings.Count(s, "u")
		count += strings.Count(s, "A")
		count += strings.Count(s, "E")
		count += strings.Count(s, "I")
		count += strings.Count(s, "O")
		count += strings.Count(s, "U")

		return count
	}


## Esercizio 1 - Capitalize
*Problema*: Scrivere un programma go ``capitalize.go`` che legge in
input una stringa e restituisce la stessa stringa in cui l'inizio di
ogni parola è maiuscolo mentre il resto è minuscolo.

*Esempi di esecuzione*:

Inserire una stringa: **nESSuno sCriVe iN quESTo moDo**

Nessuno Scrive In Questo Modo

### Soluzione


	package main

	import (
		"fmt"
		"unicode"
	)

	func main() {
		var c rune
		fmt.Print("Inserire una stringa: ")

		first := true
		for fmt.Scanf("%c", &c); c != '\n'; fmt.Scanf("%c", &c) {
			if first && ! unicode.IsSpace(c) {
				first = false
				fmt.Printf("%c", unicode.ToUpper(c))
				continue
			} else if unicode.IsSpace(c) {
				first = true
			}
			fmt.Printf("%c", unicode.ToLower(c))
		}
		fmt.Println()

		return
	}

## Esercizio 2 - Database
*Problema*: Scrivere un programma go ``database.go`` che legge in
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

	package main

	import "fmt"

	func main() {
		var input string
		fmt.Print("Inserire una riga del database: ")
		fmt.Scan(&input)
		database(input)
	}

	// name;surname;email;mark
	func database(s string) {
		var index int
		var currentField string

		for _, c := range  s {
			if c != ';' {
				currentField += string(c)
			} else {
				switch(index){
				case 0:
					fmt.Println("Nome: ", currentField)
				case 1:
					fmt.Println("Cognome: ", currentField)
				case 2:
					fmt.Println("Email: ", currentField)
				default:
				}
				currentField = ""
				index++
			}
		}
		fmt.Println("Ultimo voto: ", currentField)
		return
	}


## Esercizio 3 - Cripta password

*Problema*: Scrivere un programma ``cripta_password.go`` che legge in
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

	package main

	import "fmt"

	func main() {
		var c rune
		for fmt.Scanf("%c", &c); c != '\n'; fmt.Scanf("%c", &c) {
			fmt.Print(password(c))
		}
		return
	}

	func password(c rune) string {
		switch(c){
		case 'a':
			return "@"
		case 'e':
			return "3"
		case 'o':
			return "0"
		case 's':
			return "$"
		case 'i':
			return "1"
		case 'E':
			return "€"
		case 'L':
			return "£"
		}
		return string(c)
	}
