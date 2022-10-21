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
impostazioni
    la1=.\myexe.exe arg1 arg2 arg3
    la2=.\myexe.exe arg1 arg2 arg3
    la3=.\myexe.exe arg1 arg2 arg3

# profilo regExp
+[a-zA-Z0-9-_. \t]*Visual Studio Code
    la1=.\myexe.exe arg1 arg2 arg3
    la2=.\myexe.exe arg1 arg2 arg3
    la3=.\myexe.exe arg1 arg2 arg3

# profilo className
@Notepad
    la1=.\myexe.exe arg1 arg2 arg3
    la2=.\myexe.exe arg1 arg2 arg3
    la3=.\myexe.exe arg1 arg2 arg3

# profilo di DEFAULT
DEFAULT
    la1=info
    la2=info
    la3=info
```

Installazione
==========
per comodità si consiglia di creare dei collegamenti che avviano il programma con gli argomenti da linea di comando.
1. Creare una cartella contenente il programma, ed il file <code>config.txt</code>.
2. creare un collegamento al programma
3. entrare nelle proprietà, sotto il tab <code>collegamento</code>
4. modificare la voce <code>Destinazione</code> aggiungendo gli argomenti desidearati.



