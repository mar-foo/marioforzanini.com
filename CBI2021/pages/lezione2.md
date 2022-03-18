# Il web
> Insieme di macchine distribuite su un'area geografica
> indefinitamente ampia, interconnettibili mediante un sistema di rete
> che crea i canali necessari perchè le macchine comunichino tra loro

**Internet** è il sistema di rete prevalente al giorno d'oggi.

# Struttura

**Grafo**: insieme di nodi e collegamenti tra i nodi. Alle estremità del
grafo ci sono gli **host** (utenti), nel *mezzo* ci sono i **router**.
I router sono nodi dedicati alla sola funzionalità di rete. I router
conoscono tutta la rete, devono saper **instradare** i messaggi sulla
rete (routing).

## Velocità della rete

La velocità della rete dipende da molti fattori:

* Architettura dei router: sono in grado di gestire molte
connessioni allo stesso tempo (memoria)? quanto sono veloci nel
farlo (potenza computazionale)?
* Velocità dei link: di che materiale è fatto il collegamento tra
due nodi? Fibra ottica, wireless&#x2026;

## Unità dati

Ovviamente visto che sulla rete viaggiano grandi quantità di
informazioni bisogna trovare un modo di trasferire l'informazione
in modo veloce, efficace, affidabile e leggero. I router non
possono gestire una quantità infinita di dati allo stesso tempo.
L'informazione viene perciò suddivisa in **pacchetti**: questo
permette di ridurre il carico sui router. Il prezzo da pagare è che
ogni pacchetto dovrà contenere alcune informazioni che permettono
di essere spedito alla destinazione giusta (come quando si spedisce
una lettera), questo aggiunge informazione (e quindi peso) ad ogni
dato che viene trasportato sulla rete (si parla di decine di
byte per pacchetto).

## Affidabilità

Per garantire che l'informazione arrivi in correttamente occorre
che ci sia una verifica da parte dei due nodi che la ricevono. Per
questo motivo il punto di partenza ed il punto di arrivo si
scambiano dei messaggi per assicurarsi che tutti i pacchetti siano
stati consegnati intatti e nell'ordine corretto. Questo
complessivamente sacrifica un po' di tempo, quello necessario alla
consegna dei messaggi, ma rende più affidabili le comunicazioni.

# Navigazione in rete

La navigazione nel **World Wide Web** ("rete globale") generalmente
avviene attraverso un **browser**: Firefox, Chrome, Brave, Vivaldi,
Opera&#x2026;

## DNS

Ogni indirizzo web ha associato un indirizzo, detto **indirizzo IP**,
formato da 4 numeri separati da un punto; questo indirizzo
identifica univocamente un nodo della rete. Una parte del
connettersi alla rete sta nel tradurre gli indirizzi web (come li
conosciamo) in indirizzi IP (come li conosce la rete), questa parte
del processo è affidata al protocollo **DNS** (domain name system): un
host che ha il compito di mantenere una tabella aggiornata che
associ ad ogni indirizzo come lo conosciamo noi il corrispettivo
indirizzo IP.


## URL (uniform resource locator)

Un 'link' come lo pensiamo noi assume la forma di un URL:
\[protocollo\]://\(utente@\)\[dominio\]\(:porta\)\[percorso\]\[?query\]\[#fragment\]

### Protocollo

Il protocollo è la convenzione utilizzata per trasmettere
l'informazione, nel caso del web moderno i protocolli più diffusi
(non gli unici) sono l'**http** (e la versione criptata **https**) per
le pagine web, l'**imap** il **pop** e l'**smtp** per le e-mail. Ogni
protocollo specifica quale porta va utilizzata per la
comunicazione e quale formato debbano assumere i dati comunicati.

### \(utente@\)\[dominio\]\(:porta\)

Quello che noi comunemente chiamiamo indirizzo: il 'nome' del sito
a cui vogliamo accedere (google.com, youtube.com)&#x2026;
Notare che quello che noi comunemente chiamiamo indirizzo è in
realtà soltanto il dominio, al dominio si può poi aggiungere
informazioni riguardo all'utente con cui si vuole accedere ed alla
porta attraverso la quale si vuole accedere.

### Percorso

Il file a cui si vuole accedere sulla macchina il cui indirizzo è
dato dal dominio. In particolare **\[#fragment\]** indica quale
sezione della pagina si sta visualizzando.

### \[?query\]

Spiegheremo più avanti.

## Server

Per accedere ad una pagina web occorre che da qualche parte nel
mondo ci sia un computer acceso in grado di spedirla a chi la
richiede. L'atto di essere connessi sempre in ascolto alla
rete si dice **fare da server web**, in particolare per i computer il
cui scopo principale è fornire questa funzione si dicono server
web. Esistono molti tipi di server.

## Pagine web

Ogni pagina web è un file su un server. Ovviamente esistono diversi
tipi di file, diversi formati, non possiamo permettere che ognuno
carichi il formato che vuole altrimenti ci sarebbe incompatibilità
e servirebbero molti modi diversi per leggere le pagine web. La
scelta del formato è ricaduta su un linguaggio per scrivere
documenti chiamato HTML. L'HTML è un formato di solo testo che si
presenta così:


	<DOCTYPE! html>
		<html>
			<body>
				<h1>Una pagina web</h1>
				<p>Test.</p>
			</body>
		</html>

I browser che utilizziamo per navigare il web sono in grado di
formattare le pagine seguendo le direttive di questo linguaggio.

## Contenuti dinamici

Il problema di questo tipo di linguaggio è che non permette di
mostrare pagine create dinamicamente sulla base di alcune
informazioni che variano nel tempo. Per esempio: la pagina degli
avvisi della scuola ha sempre la stessa forma, ma i contenuti sono
diversi in base a quando la guardate; la pagina della vostra
ricerca google ha sempre la stessa forma ma i contenuti sono
diversi in base alla vostra richiesta. Per questo motivo nella
specifica del protocollo si è pensato di creare delle pagine web
speciali in cui si potesse generare automaticamente il contenuto
che viene mostrato all'utente. Alcune pagine speciali sono in grado
di eseguire dei programmi e utilizzare i risultati del programma
come contenuto della pagina. La sezione dell'url che ha la forma
**?query** serve per dare degli argomenti ai programmi che generano
la pagina. Inizialmente i programmi potevano girare solo sul
server.

## Cookie

I cookie sono dei file in cui i programmi che
generano le pagine web possono mantenere informazioni persistenti. Per esempio: il
contenuto del carrello se state comprando dei prodotti, le
informazioni riguardanti l'utente con cui avete fatto il login se
state guardando le email&#x2026; tutte queste informazioni è necessario
che sussistano per tutta la sessione in cui voi vi trovate nel sito
e vengono perciò conservati in file sul vostro computer cui i
programmi del sito possono accedere. Al giorno d'oggi se ne sente
parlare ovunque e tutti i siti pensano bene di abusarne (e sono stati
costretti per legge dall'unione europea a chiedere il consenso per
il loro utilizzo, ne parleremo).

## Javascript

Ad un certo punto della storia di internet si è pensato che l'HTML
fosse un linguaggio troppo povero per quelle che erano le esigenze
del mercato; a questo punto è stata ampliata la specifica del
protocollo aggiungendo la possibilità di utilizzare il Javascript
per generare il contenuto delle pagine. La differenza tra
Javascript e i programmi di cui parlavo prima è che i programmi in
javascript sono inclusi nel codice HTML della pagina (non lo
generano loro), inoltre quando il browser riconosce le sezioni della
pagina in cui c'è un programma scritto in questo linguaggio, esegue
il programma sul computer dell'utente (non sul server!). I
programmi scritti con questo scopo aggiornano in tempo reale la
pagina e creano quegli effetti ormai tanto abusati al giorno
d'oggi: menù a scomparsa, pubblicità fluttuanti, animazioni&#x2026;

## Pericoli ed attenzioni

L'idea di un contenuto dinamico in sè non è malvagia, ma può
essere abusata in molti modi. Basta pensare che il programma viene
eseguito sul computer dell'utente che sta navigando il web per
capire che si aggiungono una serie di problemi di sicurezza che
prima non erano nemmeno concepibili. Quando parleremo di linguaggi
di programmazione parleremo anche del perchè in molti considerano
il javascript come un linguaggio bruttissimo.
Al giorno d'oggi il javascript ed i cookie vengono utilizzati
principalmente per raccogliere informazioni sull'utente che sta
navigando, in modo da poter tracciare ogni suo minimo pensiero.
