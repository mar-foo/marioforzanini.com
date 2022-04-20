# Ripasso
La scorsa lezione abbiamo visto le slice: array con dimensioni
variabili. Come già accennato, l'utilizzo delle slice è predominante
rispetto a quello dei semplici array, poichè la generalità dei
problemi richiede di poter definire metodi che funzionino
indipendentemente dal numero di dati che occorre processare. Per
esempio il procedimento del calcolo della media è univoco e dipende in
maniera semplice dal numero di dati in input, ora siamo in grado di
scrivere una funzione che tratti il caso più generale senza
preoccuparsi della taglia dell'input:

	func media(voti []float64) float64 {
		var sum float64
		for voto := range voti {
			sum += voto
		}
		
		return sum / float64(len(voti))
	}

# Lezione finale: un esercizio interessante
Oggi vorrei terminare il corso con un esercizio che riassuma e
sintetizzi il percorso di quest anno, vi invogli a programmare in
futuro e vi costringa ad approfondire alcuni aspetti di Go che non
abbiamo ancora toccato.

Innanzitutto: l'esercizio non è pensato per essere scritto in un'ora e
mezza di lezione, ma alcune delle richieste dovrebbero risultarvi
familiari dal momento che nella parte finale del corso vi ho proposto
molti esercizi *propedeutici* in cui ne era richiesta una soluzione.
Siccome la prima regola (non scritta) della programmazione è che è
lecito copiare ed incollare, l'unico caveat è che si comprenda ciò che
viene copiato, siete autorizzati ad *ispirarvi* agli esercizi passati.

# Un server per il calcolo della media dei voti

*Problema*: Un utente del vostro sito necessita di calcolare la media
dei suoi voti, divisi per materie, ed è disposto ad inserirli uno ad
uno purchè vengano salvati per la prossima sessione. Scrivete un
programma Go che legga un file nel formato specificato all'esercizio 0
[dell'ultima lezione](programmazione13.html) e presenti all'utente la
risposta la media dei suoi voti.

## Note
### Html
Quelli di voi che hanno partecipato al primo corso sanno che un sito
web è scritto in un linguaggio che si chiama **HTML**, la cui sintassi
permette di specificare gli elementi della pagina web.

### Server web
Non mi aspetto che voi in questa ora e mezza riusciate a spulciare la
documentazione di Go ed a capire come scrivere un server web che
calcoli la media dei voti. Per questo motivo vi metto a disposizione
una certa quantità di materiale che si occupa di questi problemi.
Potete utilizzare il mio codice come infrastruttura del vostro
programma, ma vi invito a leggerlo una volta finito l'esercizio per
cercare di capire come Go si interfacci con la rete e con la creazione
di un server web (ricordate che il linguaggio è nato a Google, dove
questo è uno dei problemi più comuni da risolvere).

Il materiale è a disposizione come archivio zip [sul mio
sito](/files/cbi2021_finale.zip). La struttura è la seguente:

* Un pacchetto chiamato ``cbiHtml`` che contiene una serie di funzioni
  per scrivere **html**. Potete accedere alla documentazione con il
  comando ``go doc``
* Un file ``main.go`` che contiene le istruzioni per scrivere il main
  e le funzioni necessarie al funzionamento del programma. Potete
  scrivere tutte le funzioni che volete, ma alcune sono necessarie per
  il funzionamento richiesto: è importante che non cambiate i nomi di
  queste funzioni.

