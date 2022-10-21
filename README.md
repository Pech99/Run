About Run
==========

Il programma ha un funzionamento semplice, una volta avviato accede al titolo della
finestra in primo piano e procede al confronto tra i nomi dei diversi profili.
Trovata una corrispondenza esegue l'azione corrispondente alla lable passata da linea di comando.

I profili vanno aggiunti nel file <code>config.txt</code>

Nel file è possibile esprimere commenti con un <code>#</code> o un <code>;</code>
all'inizio della riga.

I profili sono formati da un nome e da delle opzioni indentate da del white space.
Ogni azione ha una lable seguita da <code>=</code> e da un comando.
Il comando <code>info</code> visualizza una finestra contenente le informazioni delle finestra in primo piano.

esistono 3 diversi profili:

1. profilo semplice (viene fatto il match diretto tra il nome ed il titolo della finestra)
2. profilo regExp (il match viene fatto tramite corrispndenza con espressioni regolai)
   il nome inizia con il <code>+</code> (questo carattere non fa parte dell'espressione regolare)
3. profilo di <code>DEFAULT</code> (profilo applicato nel caso in cui non vi sia una corrsopndenza)

config.txt Exemple
==========
```txt
# profilo semplice
impostazioni
    la1=info
    la2=info
    la3=info

# profilo regExp
+[a-zA-Z0-9-_. \t]*Visual Studio Code
    la1=info
    la2=info
    la3=info

# profilo di DEFAULT
DEFAULT
    la1=info
    la2=info
    la3=info
```