# Sicurezza e privacy online

Navigare online ha molti vantaggi e rende la nostra vita più semplice
e conveniente, parallelamente però si palesano anche dei problemi
nuovi che ci mettono in pericolo sotto punti di vista inimmaginabili
prima di Internet.

In particolare due sono i fronti, spesso legati tra loro, su cui
occorre prestare attenzione online: sicurezza e privacy.

## Sicurezza

Il web porta con sè delle grandi minacce alla sicurezza dei suoi
utenti: furto di identità, furto di denaro&#x2026; Ogni dispositivo che
venga connesso alla rete diventa un possibile **vettore** di attacchi.
Alcuni di questi **exploit** sono indipendenti dall'esistenza della
rete, ma vengono resi più seri nel momento in cui sono calati in un
contesto di interconnessione.

Non è difficile immaginare cosa succederebbe se chiunque avesse
accesso fisico al vostro telefono cellulare potesse sbloccarlo e
mandare messaggi a vostro nome; immaginate se queste stesse
possibilità fossero aperte a tutti i dispositivi che sono in grado di
collegarsi via rete al vostro cellulare: il rischio dell'attacco
sarebbe tanto maggiore tante più persone sono potenzialmente in grado
di compierlo.

### Prevenzione

* Non utilizzare il web

Il metodo di prevenzione più efficace è anche il più semplice: non
utilizzare la rete se non è strettamente necessario. Disconnettere (o
spegnere) il proprio dispositivo quando non è utile che sia connesso.

Al giorno d'oggi purtroppo la maggior parte delle persone si è
abituata ad essere perennemente connessa, dal punto di vista che
stiamo analizzando questo significa essere perennemente esposti.

* Ovvietà

Alcuni metodi di prevenzione sono ovvi e non penso di doverli
spiegare, mi limiterò a nominarli per completezza:

0. Non viaggare su siti strani
0. Non scaricare programmi da internet (tranne i programmi che ho
scritto io)
0. Non copiare ed incollare comandi dal web sul prompt
0. Non cliccare su link nelle email (o in altri tipi di messaggi),
sempre copiare ed incollare il link
0. Non aprire gli allegati da indirizzi email sconosciuti

* Password

Le password sono un mezzo di prevenzione di cui avete il completo
controllo: non sprecatelo. Purtroppo non si presta abbastanza
attenzione alla questione, ma ci sono diversi criteri che
andrebbero usati quando si crea una password:

0. Non usare la stessa password per 2 (o più) cose
0. Non usare una password che consista di 1/2 parole del dizionario
0. Non usare parole chiave che sono facilmente riconducibili a voi, come il nome del vostro cane o la vostra data di nascita
0. Evitare di fare un mix delle cose precedenti. 'password1234' è una password schifosa, 'p@ssword1234' è altrettanto schifosa (no, non siete i primi che hanno pensato di sostituire le lettere di una parola con simboli o numeri che sembrano simili)

Come gestire tutte queste password brutte e diverse? Avere una
password diversa per ogni servizio, non correlata alle altre e
piena di simboli può diventare difficile da controllare. La
soluzione è usare un **password manager**: un programma che gestisca
le password per voi. L'idea di programmi di questo tipo è quella
di avere una sola password che permetta di decriptare tutte le
altre che vengono conservate nel 'portachiavi'. In questo modo
occorre ricordare una sola password, ma si possono generare
password arbitrariamente difficili che vengono gestite
automaticamente. Alcuni consigli [KeepassXC](https://keepassxc.org), 
[Bitwarden](https://bitwarden.com), [Pass](https://passwordstore.org).

## Alcuni esempi di attacco

* __Phishing__:
Un attacco di **phishing** ha lo scopo di ottenere le
credenziali della vittima per l'accesso ad un servizio online. Il
tipico svolgimento di un attacco coinvolge una mail con il link
alla copia di un sito (Facebook, Instagram, Google&#x2026;) ed un
pretesto per convincere l'utente ad inserire le proprie credenziali
su questa pagina
* __DDOS__:
Un attacco **Distributed Denial Of Service** ha lo scopo di bloccare la
connessione ad internet della vittima intasando la sua rete di richieste.
* __Ransomware__:
Un attacco **ransomware** consiste nel convincere la vittima ad
eseguire un programma che cripta progressivamente tutti i file
sull'hard disk e ricatta la vittima proponendo la password per decriptarli
* __Dizionario__:
Questo tipo di attacchi è molto semplice: con un dizionario delle
password più usate si cerca di accedere al computer della vittima,
provando una password alla volta dalla lista.

## Privacy

Strettamente legato alla sicurezza degli utenti si presenta il
problema della privacy: rivelare dati su di sè apre nuove possibilità
a chi cerchi di attaccarvi.

### Segretezza vs Privatezza

Al giorno d'oggi la privacy viene spesso confusa con la segretezza e
viene per questo sminuita con argomenti come "Perchè dovrei
preoccuparmi della privacy? Non faccio nulla di male!". Il problema
con questo tipo di pensiero è che confonde il concetto di segretezza
con il concetto di privacy: nessuno di noi chiude la porta quando va
in bagno perchè deve fare qualcosa di segreto, o persino qualcosa di
male.

### Tutela della privacy

#### Cosa diciamo di noi

Le informazioni più facili da raccogliere su di noi sono quelle che
offriamo spontaneamente: foto postate sui social media, messaggi
postati online&#x2026; non dovrebbero mai contenere informazioni sensibili
che non vorreste dire a vostra madre. Questo metodo è sufficientemente
efficace, credo. Nessuno di voi posterebbe mai una fotografia della
password del suo computer, allora perchè postare il vostro indirizzo
di casa, o altre informazioni personali? Tutto ciò che va su internet
ci resta per sempre.

#### La situazione del web oggi

Navigare online oggi comporta una sfida continua alla tutela della
privacy: la maggior parte dei siti implementa dei metodi per
tracciare i propri utenti, raccogliere dati su di loro e
rivenderli al miglior offerente. La struttura del web moderno
aiuta in questo mettendo a disposizione JavaScript e Cookie senza
limitazioni di utilizzo.

Alcuni giganti del web sono più responsabili di altri, Google,
Facebook, Amazon&#x2026; sono i primi che non rispettano la privacy dei
loro utenti e obbligano coloro che vogliono collaborare a fare lo
stesso. Per esempio tutti i siti che contengono un tasto
'Condividilo su Facebook' permettono a Facebook di mantenere un
profilo dettagliato di tutti coloro che visitano il sito (anche
chi non avesse un account facebook).

Va detto che tutti questi **database** di dati dovrebbero essere
*anonimi*, ma degli studi dimostrano che vista la grande quantità
di dati raccolti non è difficile correlare certi dataset con
persone specifiche.

#### Opzioni per difendere la propria privacy
* Navigazione in incognito

La navigazione in incognito è il più grande mito che va sfatato
nel contesto della tutela della privacy. La navigazione in
incognito cancella la cronologia dei siti visitati sul vostro
browser, ma mantiene la lista di cookies e dati dei siti che
avete visitato; questo significa che dal punto di vista di
Google, Facebook e Amazon è come se non aveste cancellato nulla
perchè il mezzo attraverso cui siete riconosciuti sono i cookie.
Ciò che bisogna fare se si vogliono cancellare tutti i dati di
navigazione è istruire il proprio browser in modo che li elimini
ogni volta che viene chiuso il programma.

**Chrome**

Navigare alla pagina 'chrome://settings/content/cookies' ed abilitare l'opzione che dice di tenere i dati locali fino all'uscita

**Firefox**

Navigare alla pagina 'about:preferences#privacy' e selezionare l'opzione di cancellare cookies e dati dei siti quando Firefox viene chiuso.

* Https

Un altro modo per tutelare la propria privacy è quello di
sfruttare l'https che stabilisce connessioni criptate verso il
server web con cui comunicate. Nonostante la connessione sia
criptata, per entrare sulla rete dovete passare attraverso una
richiesta DNS al vostro router (che generalmente la inoltra ai
server DNS del vostro ISP), questo significa che la lista di siti
che avete visitato è nota al vostro provider internet.


* VPN

Per ovviare al fatto che il vostro provider internet possegga una
lista dei siti da voi visitati si può ricorrere all'uso di una
VPN: un meccanismo che rimbalza il vostro segnale attraverso
alcuni server che fanno la richiesta per voi. Anche in questo
caso però la maggior parte dei servizi che offrono questa opzione
sono in grado di ricondurre alcune informazioni al vostro account
e siete costretti a fidarvi di loro almeno tanto quanto vi fidate
del vostro ISP.

* Tor

La soluzione più completa dal punto di vista della privacy è
sicuramente usare Tor, ogni vostra richiesta DNS viene risolta
sul vostro computer e tutto il traffico viene nascosto attraverso
l'**onion routing**. Ci sono alcuni svantaggi nell'utilizzo di Tor:

0. La rete è molto lenta
0. Alcuni siti bloccano gli indirizzi Tor
0. Potrebbe dare un senso di falsa sicurezza, non essendo un arma infallibile va utilizzato bene
