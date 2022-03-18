# Criptovalute

## Blockchain

La blockchain è un database distribuito e pubblico

* Distribuito

Le informazioni contenute nella blockchain vengono comunicate a
tutti i nodi nella rete, non c'è un database centrale. Ciò che prova
l'autenticità dei dati è il confronto dei blocchi posseduti da ogni
nodo separatamente

* Pubblico

Chiunque può esaminare la blockchain e cambiarla

La creazione di nuova informazione da aggiungere alla blockchain
consiste nel creare un hash dell'informazione che deve essere
conservata. Un algoritmo di hashing produce una stringa che identifica
univocamente i dati in ingresso, in particolare l'hash di un nodo
della blockchain contiene sempre tra i dati in ingresso l'hash del
blocco precedente. In questo modo ogni blocco si riferisce ai
precedenti in una catena che procede dal **blocco di genesi**.

Poiché l'informazione di ciascun blocco è conservata nei blocchi
successivi, cambiare un anello della catena richiede una grande
potenza computazionale per cambiare anche tutti i successivi.

## Fork

Una fork (**biforcazione**) è una ramificazione: un punto in cui la
blockchain si divide in due rami separati che crescono in maniera
indipendente. Normalmente questo non dovrebbe accadere, ma siccome
l'ordine dei blocchi dipende solo dal momento in cui sono creati, può
capitare che due blocchi siano aggiunti allo stesso tempo; in questi
casi la blockchain si biforca e ogni nuovo blocco viene aggiunto ad
entrambe i rami. Quello che accade molto spesso è che gli utenti
decidono di mantenere viva una ramificazione che diventa la più lunga
e perciò considerata la *vera* blockchain, a questa verranno aggiunte
tutte le transazioni che stanno sulla catena abbandonata (**orphan
chain**). In altre occasioni la divisione è permanente e crea due
blockchain separate ed indipendenti.


## Un sistema zero trust

La rivoluzione della tecnologia blockchain sta nella possibilità di
creare un sistema distribuito in cui ciascun utente agisce per sè e il
buon funzionamento del sistema è dovuto all'architettura della
blockchain stessa. Non occorre avere fiducia negli utenti della rete
poichè il funzionamento della rete stessa rende molto difficile (anche
se non impossibile) ingannare gli altri e cambiare le informazioni nel
database in proprio favore. Dal punto di vista delle valute questo è
un cambiamento epocale: non occorre più affidarsi allo stato, o alle
banche, perchè il valore dei propri soldi sia ben definito: il valore
delle criptovalute è deciso dal sistema e dalla richiesta di moneta,
non è necessario fidarsi di un ente centrale che controlla tutto.

## Valute

### Bitcoin

Nel caso specifico di Bitcoin la blockchain viene utilizzata per
conservare informazioni riguardo alle transazioni effettuate sulla
rete.

Completare una transazione significa fare una richiesta di aggiunta
alla blockchain, alcuni utenti della rete (i **miner**) si occupano di
risolvere i problemi crittografici necessari per aggiungere un blocco
alla catena. Il primo blocco prodotto viene aggiunto alla
blockchain ed il miner che lo ha prodotto viene ricompensato con delle
monete. La grande competizione per creare nuovi blocchi (che vengono
aggiunti con una cadenza fissata di 1 blocco/minuto, ed hanno
grandezza massima fissata) comporta grandi aumenti nelle tasse
necessarie a creare una transazione: per invogliare i miner ad
occuparsi della propria transazione prima delle altre occorre pagare
delle tasse più alte.

#### Portafoglio

Un portafoglio di bitcoin consiste in una coppia di chiavi: una chiave
privata (posseduta dal proprietario del portafoglio) ed una chiave
pubblica conservata nella blockchain. Soltanto il proprietario della
chiave privata può utilizzare i Bitcoin conservati nel portafoglio,
al contrario la quantità di monete all'interno di un portafoglio è
scritta nella blockchain ed è di dominio pubblico.

### Monero

Una moneta interessante che è nata negli ultimi anni è Monero. Lo
scopo di Monero è mitigare il problema dell'eccesiva trasparenza di
Bitcoin: la tecnologia della blockchain funziona bene, ma
l'implementazione di Bitcoin permette ad un qualsiasi utente della
rete di risalire al saldo di ogni portafoglio. Con Monero ogni
transazione subisce dei cambiamenti atti a nascondere al resto della
rete l'identità delle parti in gioco.

### Ethereum

Ethereum viene pensato come una piattaforma distribuita che permette a
chiunque di creare, conservare ed eseguire delle applicazioni
distribuite (**DApps**) basate su contratti. Anche in questa valuta si
utilizza la tecnologia della blockchain per mantenere lo stato globale
della rete, ma si concentra sul creare una piattaforma in cui creare
applicazioni che sfruttino questa architettura per offrire servizi
nuovi e distribuiti.

### ecc.

C'è una grande quantità di criptovalute che nascono ogni giorno,
poiché il successo di Bitcoin porta molti a tentare di ripetere
l'impresa. Sarebbe impossibile parlare di tutte le monete esistenti.

## Valore

Perchè una criptovaluta dovrebbe avere valore? Come dare valore ad una
entità solamente digitale?

Ciò che conferisce valore alla moneta è la scarsità: le criptovalute
sono il primo esempio di entità digitale non facilmente copiabile.
Occorre un grande lavoro computazionale per minare Bitcoin, il verbo
stesso utilizzato per descrivere questa operazione si riferisce
all'analogia tra Bitcoin e oro: anche l'oro assume un valore sulla
base della sua scarsità e della difficoltà di ottenerne altro (occorre
appunto estrarlo, minarlo). Questa scarsità è dovuta al funzionamento
della blockchain.

## Risorse

* [Bitcoin wiki - Blockchain Explanation](https://en.bitcoinwiki.org/wiki/Blockchain)
