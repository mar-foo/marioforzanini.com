# Programmazione
I programmi sono quello che fa funzionare l'hardware, le istruzioni che la macchina deve eseguire. Poiché la CPU comprende una serie di istruzioni ben precise, sotto forma di codice binario, e diverse per le diverse architetture, il modo di scrivere programmi se è evoluto nel tempo.

* Inizialmente i programmi venivano scritti in linguaggio macchina ed i bit venivano attivati a mano per caricare il programma
* Successivamente fu inventato l'assembly: un linguaggio che permette di esprimere le istruzioni che la macchina sa comprendere in un linguaggio umano che viene poi **compilato** da un *assembler*. __Compilare__ significa trasformare il programma scritto in assembly in una sequenza di bit che hanno senso per la macchina
* Vengono inventati linguaggi __strutturati__ (ad alto livello) che permettono di scrivere il programma in un linguaggio meno vicino alla macchina ma che può essere tradotto per più architetture differenti

**Go** appartiene all'ultima categoria

## Compilazione
I linguaggi "ad alto livello" si suddividono poi in ulteriori classi (procedurali, funzionali, staticamente tipizzati...). In particolare:

* Un linguaggio si dice compilato se il codice sorgente per essere eseguito necessita di venire trasformato in una serie di istruzioni binarie
* Un linguaggio si dice interpretato se viene eseguito da un programma

# Golang
I linguaggio nasce a Google dalle idea di Ken Thompson, Rob Pike e
Robert Griesemer. Due membri del gruppo che ha lavorato su Unix, Plan9
ed Inferno. Ken Thompson è uno dei padri del C (linguaggio nato negli
anni '70 ed utilizzato ancora oggi per applicazioni di ogni tipo nel
campo della progettazione di sistemi). Go viene concepito inizialmente
nel 2007, viene rilasciato pubblicamente nel 2009. ('Il C del XXI
secolo')

## Idee
Semplicità, leggibilità, velocità di compilazione e scrittura.

## Installazione
Seguire le istruzioni del [sito ufficiale di Go](https://go.dev). Test
dell'installazione: creare un file 'test.go' con i seguenti
contenuti

	package main
	
	import "fmt"
	
	func main() {
		fmt.Println("Hello world!")
	}

Aprite un <b>prompt dei comandi</b> nella stessa cartella e scrivete:

	go build test.go
	test.exe

Dovrebbe apparire a schermo "Hello world!"
