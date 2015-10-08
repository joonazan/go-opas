# go-opas
Suomenkielinen opas Go-kieleen aloittelijoille. Epälineaarinen, muistuttaa teknologiapuuta. Pyrkii kattamaan kaikki kielen ominaisuudet.

Opas pyrkii olemaan kertomatta epäolennaisia yksityiskohtia heti, kuitenkaan opettamatta mitään väärin yksinkertaisuuden vuoksi.

## Yksinkertainen Go-ohjelma

```Go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

```

Tietokoneet osaavat _ajaa_ yksinkertaisista käskyistä koostuvia ohjelmia. Näitä käskyjä kutsutaan usein _konekieleksi_. Konekielisten käskyjen luetteleminen on kuitenkin niin työlästä ja vaikeaa, että jo 1950-luvulla luotiin ensimmäiset _ohjelmointikielet_.

Nykyään ohjelmia tehdään kirjoittamalla _lähdekoodia_ jollain ohjelmointikielellä. Koodi voidaan muuttaa konekieleksi _kääntäjällä_ tai sitten voidaan ohjelmoida _tulkki_, joka tekee lähdekoodin kuvaamat asiat. Tähän oppaaseen on valittu kieli nimeltä *Go* sen yksinkertaisuuden ja hyödyllisyyden vuoksi.

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
