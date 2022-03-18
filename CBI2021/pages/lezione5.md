# Premessa

## La struttura dei file

L'ultima volta ho avuto l'impressione che ci fosse un po'
di confusione riguardo alla struttura di file, cartelle ecc. Colpa
mia che l'ho dato per scontato.
I file nel computer sono organizzati in una struttura che è
assimilabile a quella di un albero. A partire dalla radice ogni
cartella è come il nodo di un albero, sotto ogni cartella si
diramano i file (le foglie dell'albero) e le cartelle in essa
contenute.
Per riferirsi ad un file si possono usare due modi:

* Dare un percorso che parte dalla radice dell'albero e percorre
tutti i nodi necessari per raggiungere il file (percorso
assoluto)
* Dare un percorso verso il file a partire dalla cartella in cui ci
si trova (percorso relativo). Per riferirsi alla cartella
precedente nell'albero si usa il simbolo '..'

Dal punto di vista del gestore grafico di documenti: dare un
percorso assoluto significa tornare sempre indietro finchè si può e
seguire tutto il percorso verso il file, dare un percorso relativo
significa procedere a ritroso fino alla prima cartella in comune e
poi procedere con il percorso verso il file.

## Estensioni dei file

La seconda impressione che ho avuto è stata che il concetto di
formato di un file e relativa estensione fosse sconosciuto. Diversi
tipi di informazione richiedono diversi modi di essere salvati in un
file, per esempio: per i video ci sono **mp4**, **mkv**&#x2026; per le
immagini **png**, **jpg**, **gif**&#x2026; Per convenzione si indica un file
che contiene un certo tipo di dato aggiungendo l'estensione relativa
al suo formato alla fine del nome. L'ultima volta questo forse era
risultato un po' confuso perchè di default il file manager di
windows nascondeva l'estensione, ma il nome del file **comprende
l'estensione**! Il file "immagine.jpeg", il file "immagine.png" ed il
file "immagine" sono diversi, inoltre windows si basa
sull'estensione per capire di che tipo di file si tratti e per
proporre un programma adatto. Tutti i file scritti in linguaggio di
markup, oppure in un qualunque linguaggio di programmazione sono
file di testo semplice, ma l'estensione che si usa dipende dal
linguaggio scelto: **.go** per i file go, **.html** per i file HTML,
**.cpp** per i file C++&#x2026;

## Server

Dopo la premessa riguardo alla struttura del **filesystem** forse è
più semplice capire cosa significa che la struttura di un sito web è
soltanto una cartella sul server. Nel momento in cui guardo un sito
nella barra dell'indirizzo appare un percorso relativo che mi porta
dalla cartella iniziale fino al file che sto guardando. Il sito può
essere organizzato con una struttura più o meno complicata di
cartelle e generalmente in ogni cartella ci sarà un file chiamato
index.html (per convenzione). Ogni link che punti a pagine
all'interno del sito **deve** specificare un percorso **a partire dalla
cartella in cui si trova il file** dove c'è il link **verso il file a
cui si vuole puntare**. Per esempio se io mi trovo in una cartella
che contiene due file: 'index.html' e 'test.html' un link dal file
index.html al file test.html avrà questa faccia:

	<a href="test.html">Test</a>

poiché i file **sono nella stessa cartella**.
Al contrario se il file test.html si trovasse nella cartella
'articoli' un link da index.html a test.html avrebbe questa faccia:

	<a href="articoli/test.html">Test in un'altra cartella</a>

Mentre un link da test.html ad index.html sarebbe così:


	<a href="../index.html">Indice</a>

Lo stesso va fatto quando si aggiungono immagini, video
(controllando che l'estensione sia corretta)&#x2026;


# Nuovi concetti di rete

## Sottorete

Come abbiamo detto nella prima lezione sul web, la rete è
soltanto un grafico con tanti nodi collegati tra loro e
identificati con degli indirizzi IP. Siccome gli indirizzi IP non
sono infiniti occorre utilizzarli in modo intelligente per servire
tutti gli utenti della rete senza duplicazioni. Per questo la rete
in realtà risulta organizzata in **sottoreti**. Un esempio sono le
reti di casa: in una casa possono esserci molti dispositivi
connessi alla rete, ma questi non hanno un indirizzo __pubblico__
ovvero questi dispositivi non sono raggiungibili dall'esterno
direttamente: quello che succede è che all'interno della rete
privata ognuno ha il suo indirizzo **privato** (per convenzione
iniziano tutti con 192.168.0.), l'unico dispositivo che ha un
indirizzo pubblico visibile a tutta la rete internet è il router; a
sua volta il router fa parte di una sottorete di router, per
esempio quella gestita da TIM. Il router che connette due sottoreti
si chiama **gateway**, ogni gateway conosce bene le sottoreti a cui è
collegato e sa come indirizzare i pacchetti all'interno di queste
sottoreti.

## Porte

Il traffico di rete in entrata può essere di diversa natura: un
computer potrebbe scaricare email, navigare su internet e
connettersi in SSH (da remoto) ad un altro; occorre un metodo per
dividere in maniera coerente il traffico in ingresso. La soluzione
adottata per questo problema è quella di definire un canale diverso
per ogni protocollo, questi canali si chiamano **porte**. Come in un
porto ci sono diversi moli, a cui attraccano diversi tipi di
imbarcazione; così il computer è in grado di ricevere traffico di
rete su diverse porte. Le porte che interessano a noi sono quelle
del traffico web che sono **80** (per l'http) e **443** (per l'https).

## Servire un sito web

Per avere un sito web occorre innanzitutto configurare il proprio
computer in modo che si metta in ascolto sulla porta corretta
(porta 80 nel nostro caso); solitamente per fare questa cosa si
installano dei programmi appositi che si chiamano **server http**,
l'unico scopo di questi programmi è quello di **servire un sito**:
rispondere ad ogni chiamata sulla porta 80 con la pagina richiesta
del vostro sito, le pagine del sito saranno cercate nella cartella
che voi indicate.

# Uniamo i puntini

Ho scritto un semplicissimo server, mi piacerebbe che ognuno
di voi scrivesse un piccolo sito fatto così:

* Una pagina di presentazione (index.html): chi siete, quanti anni
avete, che liceo frequentate &#x2026;
* Una o più pagine a vostra scelta in cui mettete quello che volete.
Queste pagine devono essere raggiungibili in qualche modo: a
partire dall'indice devono esserci tutti i link necessari per
arrivare a queste pagine

Mentre state lavorando potete attivare un programma che ho scritto
io che fa da server web, il codice è disponibile [qui](/files/server.go). Questo
programma lo trovate [sulla mia pagina web](/files/server.exe) e dovete metterlo nella
stessa cartella in cui creerete il vostro sito. Se avete capito
qualcosa di quello che abbiamo detto delle sottoreti, saprete che le
pagine web che state per scrivere saranno accessibili soltanto a
coloro che sono connessi alla rete interna della scuola a cui siete
connessi voi, quindi gli unici che possono vederle siamo noi.

## Istruzioni per attivare il programma

Per far partire il server dovete salvare il file nella cartella in
cui si trova il vostro sito, aprire quella cartella nel file
manager, schiacciare con il tasto destro e aprire un terminale
nella cartella. Nel terminale scrivete

	server.exe

per far partire il programma **server.exe**. Se qualcosa non funziona
è probabile che non siate nella cartella giusta: prestate
attenzione ad aprire il terminale dal file manager mentre siete
nella cartella del sito.

## Vedere i siti degli altri

Per vedere i siti dei vostri compagni, e perchè il vostro sito sia
visibile a loro, è necessario che diciate il vostro indirizzo ip
privato nella rete della scuola; per trovarlo aprite un prompt dei
comandi e scrivete **ipconfig**. Scrivendo nella barra degli
indirizzi del browser uno degli indirizzi dei vostri compagni
dovreste vedere apparire il loro sito.
