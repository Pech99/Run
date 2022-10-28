About Run
==========

## Generale
Il programma ha un funzionamento semplice, una volta avviato accede al titolo della
finestra in primo piano e procede al confronto tra i nomi dei diversi profili.
Trovata una corrispondenza esegue l'azione corrispondente alla lable passata da linea di comando.

### Profili e Comandi
I profili vanno aggiunti nel file <code>config.txt</code>
I profili sono formati da un nome e da delle opzioni indentate da del white space.
Ogni azione ha una lable seguita da <code>=</code> e da un comando.

esistono 4 diversi profili:

1. profilo semplice, si mecha il nome del profilo con il titolo della finestra in primo piano.
2. profilo regExp, si cerca un mech del titolo del profilo come espressione regolare nel titolo della finestra in primo piano. 
   Il nome del profilo inizia con il <code>+</code> (questo carattere non fa parte dell'espressione regolare)
3. profilo className, si mecha il nome del profilo con il nome della classe della finestra in primo piano.
  Il nome del profilo inizia con la <code>@</code> (questo carattere non fa parte dell'nome della classe del finestra in primo piano).
4. profilo di <code>DEFAULT</code>, applicato nel caso in cui non vi sia una corrsopndenza.

Il comando <code>info</code> visualizza una finestra contenente le informazioni delle finestra in primo piano.

### Commenti
Nel file è possibile esprimere commenti con un <code>#</code> o un <code>;</code>
all'inizio della riga.

## Exemple config.txt
```txt
# profilo semplice
Lettore multimediale
	sing=.\Click <MEDIA_PLAY_PAUSE>
	dopp=.\Click <MEDIA_NEXT_TRACK>
	pres=.\Click <MEDIA_PREV_TRACK>

# profilo regExp
+YouTube
	sing=.\Click <MEDIA_PLAY_PAUSE>
	dopp=.\Click <LEFT>
	pres=.\Click <F>

# profilo className
@PotPlayer64
	sing=.\Click <MEDIA_PLAY_PAUSE>
	dopp=.\Click <MEDIA_NEXT_TRACK>
	pres=.\Click <MEDIA_PREV_TRACK>

# profilo di DEFAULT
DEFAULT
	sing=.\Click <SUPER>
	dopp=info
	pres=explorer.exe
```

Installazione
==========
per comodità si consiglia di creare dei collegamenti che avviano il programma con gli argomenti da linea di comando.
1. scarica l'ultima versione [qui](https://github.com/Pech99/Run/raw/master/Run.exe)
2. Creare una cartella contenente il programma, ed il file <code>config.txt</code>.
3. creare un collegamento al programma
4. entrare nelle proprietà, sotto il tab <code>collegamento</code>
5. modificare la voce <code>Destinazione</code> aggiungendo gli argomenti desidearati.



