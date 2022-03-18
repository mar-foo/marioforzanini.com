# Fogli di calcolo

* Righe (numerate) e colonne (indicate con lettere)
* All'incrocio tra una riga ed una colonna c'è una cella
* Le celle si indicano concatenando la lettera della colonna ed il
numero della riga corrispondenti: ex: D3 AB9 K74
* Nelle celle possiamo scrivere numeri, parole, formule
* Ogni file può essere costituito da più fogli (griglie)

# Utilizzo

* I fogli di calcolo sono utili per set di dati in costante aggiornamento
* La vera potenza dei fogli di calcolo deriva dall'ordine in cui si
sanno porre le celle
* Disporre le celle nell'ordine corretto significa disporle in modo
da poter utilizzare una stessa formula sul maggior numero di celle
possibili
* Le formule servono per semplificarsi la vita


* [Raccolta dati per un esperimento](/files/Dati_Reticolo_Casa.ods)

Per le formule creare una colonna per ogni categoria di dato ed una riga per ogni membro di quella categoria.

## Esercizi

* Creare una tabella con i propri voti di scuola ed ottenerne la media
* Creare una tabella che contenga le età di ognuno. Ricavare il numero di anni vissuti in totale, il numero di giorni vissuti in totale, il numero di minuti vissuti in totale... <a href="files/eta.ods">Risultato</a>

# Formattazione

Le celle possono ereditare una serie di attributi grafici: colore, grandezza del font ... Questi attributi possono anche essere assegnati dinamicamente sulla base del risultato di formule.

## Esercizi

* Riprendere l'esempio della media e scrivere una regola per la **formattazione condizionale** che colori di rosso le celle a cui media è minore di 6
* Supponiamo che i vostri genitori vi elargiscano una paghetta e che vogliate tenere conto delle spese: creare una tabella fatta così:

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">
<colgroup>
<col>
<col>
<col>
</colgroup>
<thead>
<tr>
<th scope="col">Data</th>
<th scope="col">Spesa</th>
<th scope="col">Causale</th>
</tr>
</thead>
<tbody>
<tr>
<td>1/10/2021</td>
<td>2</td>
<td>Giornale</td>
</tr>

<tr>
<td>1/10/2021</td>
<td>15</td>
<td>Libro</td>
</tr>
</tbody>
</table>

ed un'altra così:

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">
<colgroup>
<col >
<col >
<col >
</colgroup>
<thead>
<tr>
<th scope="col">Data</th>
<th scope="col">Entrata</th>
<th scope="col">Causale</th>
</tr>
</thead>
<tbody>
<tr>
<td>1/10/2021</td>
<td>10</td>
<td>Paghetta</td>
</tr>
</tbody>
</table>

Ottenere un bilancio totale delle spese e delle entrate, colorare di rosso la casella se il bilancio è negativo. [Risultato](files/bilancio.ods)


# Grafici

A partire da tabelle di dati si possono ricavare una serie di
grafici che ne rappresentino delle qualità specifiche. Ogni tipo di
grafico ha una precisa funzione e risulta utile per osservare
qualità differenti. Ad esempio un istogramma risulta utile per
confrontare risultati che derivano dallo stesso set e confrontarne
la frequenza, un grafico a dispersione serve a confrontare
l'andamento di un particolare dato in funzione del valore di un
altro (osservare la relazione tra due dati)...

## Esercizio

A partire dal foglio con i vostri voti, se non l'avete già fatto,
dividete i voti per materia e create un istogramma che mostri la media
di ogni materia comparata con le altre. Mettete sull'asse orizzontale
la materia e su quello verticale la media.

