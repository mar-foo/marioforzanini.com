# HTML

Il linguaggio principe per scrivere siti web è l'HTML; si tratta di
un linguaggio di **markup**, ovvero un linguaggio per formattare i
documenti. Di per sè un documento HTML è un normale file di testo,
ma il browser sa interpretare la sintassi (struttura) del
linguaggio e trasformarla in elementi grafici (titoli, paragrafi,
link, immagini&#x2026;).

## Tag

L'HTML è facilmente riconoscibile perchè la struttura base del
linguaggio sono le **tag**. Una tag è un contenitore di informazione
che sintatticamente viene richiuso tra '<'. La forma è:


	<tipo opzione1="o" opzione2="a">Contenuto</tipo>

Nel primo blocco '<' indica l'apertura del contenitore, 'tipo'
indica il tipo di contenitore, 'opzione1=&#x2026;' sono le opzioni che
si possono applicare a quel tipo di contenitore (per esempio per
un'immagine si può pensare alle dimensioni, per il testo al colore,
al font, alle sottolineature&#x2026;), '>' indica la chiusura
dell'intestazione della tag e dà inizio al contenuto; ogni volta
che viene aperta una tag, questa deve anche essere chiusa '</tipo>'
chiude l'ultima tag corrispondente dello stesso tipo.

### Intestazione

All'inizio di un documento HTML si possono aggiungere delle
informazioni che non sono propriamente contenuto della pagina, ma
che possono servire al browser (**metadata**), per esempio la lingua
in cui è scritta la pagina, l'autore&#x2026; La parte più importante di
questa intestazione è la dichiarazione del tipo di documento:
l'HTML non è l'unico linguaggio accettato per scrivere siti web,
anche se è il più diffuso, quindi occorre avvisare il browser
riguardo al tipo di documento che stiamo scrivendo; per farlo
la prima riga del file deve essere del tipo:


	<DOCTYPE! html>

Dopo aver avvisato riguardo al tipo di file occorre anche
circoscrivere la sezione dove vogliamo scrivere il documento in
una tag <html>. Ciò che viene mostrato nella pagina (la sostanza
del documento) dovrà inoltre stare all'interno di una tag
'<body>'. Alla fine lo scheletro del documento sarà fatto così:


	<DOCTYPE! html>
		<html>
			<body>
	
				Contenuto qui
	
			</body>
		</html>

### Titoli

Le tag tipiche dei titoli sono numerate a partire da 1, numeri più
grandi corrispondono ad una subordinazione maggiore.


	<h1>Titolo</h1>
	<h2>Sottotitolo</h2>

### Paragrafi

I paragrafi di testo vanno inseriti in tag '<p>'.

### Link

I link a pagine sia sul proprio sito che su altri siti vanno
inseriti all'interno di tag a. Il contenuto della tag è la
descrizione del link (quella che si vede sul documento), mentre la
destinazione del link va inserita come opzione **href** in questo
modo:

	<a href="la_mia_pagina_2.html">La mia pagina 2</a>
	<a href="http://marioforzanini.com">Il mio sito</a>

### Immagini

Per inserire un'immagine all'interno della pagina si usa la tag
'<img>'. Si specifica il percorso all'immagine nell'opzione 'src':

	<img src="immagine1.jpg">
	<img src="immagini/1.jpg">

La tag *img* **non** va chiusa.

### Liste

Esistono due tipi di liste: numerate e non. Le liste non numerate
si indicano con la tag <ul>, quelle numerate con la tag <ol>.
All'interno di una tag 'lista' si devono poi aggiungere gli
elementi, ciascuno di questi va racchiuso entro una tag <li>:


	<ul>
		<li>Elemento uno</li>
		<li>Elemento 2</li>
	</ul>

Per aggiungere sottoliste è sufficiente inserire una tag <ul>
oppure <ol> al posto di un elemento della lista.

## Servire pagine web

Perchè il vostro computer possa rispondere ad altri che richiedono la
vostra pagina occorre far partire un programma (un server http),
indirizzarlo alla cartella in cui ci sono i vostri file html ed
aspettare che qualcuno si connetta. A questo punto nella rete
**interna** gli altri possono connettersi al sito digitando l'indirizzo
IP del nostro computer.

# CSS

Le pagine web scritte solo in html sono un po' noiose: non ci sono
colori e tutto viene disposto in sequenza come lo scrivete&#x2026;
Per aggiungere informazioni sulla disposizione della pagina,
piuttosto che su altri elementi stilistici si utilizza un altro
linguaggio chiamato **CSS**.
Ci sono diversi livelli a cui si può operare per aggiungere lo
stile: a livello della singola tag (e delle sottotag) oppure a
livello globale su tutto il file.

## Sintassi

La sintassi del CSS è molto semplice: si può fare riferimento ad un
tipo di tag e cambiarne gli attributi; per esempio il colore, lo
sfondo, aggiungere un riquadro attorno all'elemento ecc. I
cambiamenti vengono ereditati dalle tag figlie, a meno che non
vengano sovrascritti da ulteriori modifiche. Lo stile si presenta
con il nome della tag, ':' ed una lista di attributi separati
da ';' e racchiusi in parentesi graffe. Per esempio:

	body {
		color: black;
	}

## Operare sulle singole tag

Una delle opzioni che si possono specificare quando si apre una tag
è proprio lo stile CSS che si vuole applicare a quella tag, non è
necessario speficare il tipo di tag all'inizio del CSS.
Esempio: voglio che un link sia rosso


	<a href="example.com" style="color: red;">Link</a>

Esempio: voglio che i link siano tutti rossi:

	<body style="a {color: red; }">

## Modifiche globali

Per applicare modifiche di stile a tutto il file ci sono 3 modi,
in ordine crescente di frequenza con cui vengono usati:

* Applicare i cambiamenti alla tag <body>
* Aggiungere dentro i metadata del file le specifiche dello stile
dentro una tag <style>
* Metterli in un file separato

L'ultima opzione è in generale la più usata (ormai quasi l'unica
usata), quindi impareremo quella.
Nella testa del file aggiungere queste righe:

	<head>
		<link rel="stylesheet" type="text/css" href="style.css">
	</head>

	I contenuti del file ``style.css``
	
	a {
		color: red;
	}
	
	p {
		background: blue;
	}
	
	src {
		width: 400;
	}

# Linguaggi di markup

Esistono una serie di linguaggi di markup in circolazione che
possono essere usati per scrivere documenti. La maggior parte di
questi linguaggi, a differenza dell'HTML che viene interpretato
direttamente dal browser, devono essere compilati (trasformati in un
formato standard come PDF oppure PostScript); occorre perciò
installare dei 'compilatori' che facciano questo lavoro.
Ogni linguaggio ha diverse comodità, ed in generale viene usato per
scopi diversi, per esempio:

* Latex

Viene usato in ambito accademico per scrivere tesi e
papers, permette di scrivere equazioni con semplicità, gestisce
automaticamente i riferimenti, le bibliografie e numera
automaticamente tutto in modo dinamico. Un estensione del
linguaggio, chiamata Beamer, permette anche di scrivere delle
presentazioni in latex utilizzando tutto il potere del linguaggio
per fare delle slide

* Markdown

Molto semplice e veloce da scrivere, viene utilizzato su alcuni siti (e.g.
<a href="https://github.com">Github</a>)

* Troff, Groff, Nroff

Antico linguaggio per fare 'typesetting', pensato per
scrivere libri, papers e qualsiasi altro tipo di documento
professionale; è il più antico e permette una flessibilità
incredibile per quanto riguarda la formattazione della pagina, al
prezzo di essere un po' alieno da leggere e scrivere.

* Altri

Io vorrei farvi vedere altri linguaggi, se siete interessati, ma
penso che al giorno d'oggi non abbiate bisogno di saperli usare.
Sappiate che esistono e che sono abbastanza semplici da imparare. Il
messaggio è che esistono alternative (anche molto più comode di
Word) per scrivere documenti.

# .Docx

Se volete usare questi formati per scrivere documenti e provare un
po' a vedere se funzionano per voi, ma i vostri professori
richiedono che consegnate tutto in formato '.doc' o '.docx'
sappiate che c'è un programma che si chiama **pandoc** che compila i
documenti anche in quel formato.

