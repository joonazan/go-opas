# Alku
Tietokoneet osaavat _ajaa_ yksinkertaisista käskyistä koostuvia ohjelmia. Näitä käskyjä kutsutaan usein _konekieleksi_. Konekielisten käskyjen luetteleminen on kuitenkin työlästä, eikä lopputulos edes aina tee sitä mitä sen piti! Siksi jo 1950-luvulla luotiin ensimmäiset _ohjelmointikielet_.

Nykyään ohjelmia tehdään kirjoittamalla _lähdekoodia_ jollain ohjelmointikielellä. Koodi voidaan muuttaa konekieleksi _kääntäjällä_ tai sitten voidaan käyttää _tulkkia_, joka tekee lähdekoodin kuvaamat asiat. Tähän oppaaseen on valittu kieli nimeltä __Go__, koska se on yksinkertainen, sillä on vaikea tehdä vahingossa muuta kuin mitä halusi tehdä ja sillä ohjelmointi on niin miellyttävää, että monet käyttävät sitä.

## Hello, world!

```Go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

```

## Miten esimerkkiohjelma toimii?
```Go
func main() {
	fmt.Println("Hello, world!")
}
```
[`Println`](https://golang.org/pkg/fmt/#Println) on _funktio_, joka saa ohjelman _tulostamaan_ jotain. Se määritellään _paketissa_ [fmt](https://golang.org/pkg/fmt/). Sitä _kutsutaan_ _merkkijonolla_ `Hello, world!`.

`main` on funktio, jota kutsutaan automaattisesti kun ohjelma käynnistetään.

```Go
import "fmt"
```
[`import`](import.md) ottaa käyttöön paketin.

```Go
package main
```
`package` kertoo minkä nimiseen pakettiin tiedosto kuuluu. Jos paketin nimi on `main`, se on ohjelma, jos nimi on mikä tahansa muu, pakettia ei voi ajaa. Joka tiedoston alussa on oltava pakettimäärittely.

## Mitä värit tarkoittavat?
