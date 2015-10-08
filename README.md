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

Tietokoneet osaavat _ajaa_ yksinkertaisista käskyistä koostuvia ohjelmia. Näitä käskyjä kutsutaan usein _konekieleksi_. Konekielisten käskyjen luetteleminen on kuitenkin niin työlästä ja vaikeaa, että jo 1950-luvulla luotiin ensimmäiset ohjelmointikielet.

Nykyään ohjelmia tehdään _kirjoittamalla lähdekoodia_ jollain _ohjelmointikielellä_. Koodi voidaan muuttaa konekieleksi _kääntäjällä_ tai sitten voidaan ohjelmoida _tulkki_, joka tekee lähdekoodin kuvaamat asiat. Tähän oppaaseen on valittu kieli nimeltä *Go* sen yksinkertaisuuden ja hyödyllisyyden vuoksi.

## Miten esimerkkiohjelma toimii?
```Go
func main() {
	fmt.Println("Hello, world!")
}
```
`Println` on _funktio_, joka saa ohjelman _tulostamaan_ jotain. Se _määritellään_ kirjastossa `fmt`. Sitä _kutsutaan_ _merkkijonolla_ `Hello, world!`.

`main` on funktio, jota kutsutaan automaattisesti kun ohjelma käynnistetään.

```Go
import "fmt"
```
`import` ottaa käyttöön kirjaston. `fmt` vaaditaan, koska ohjelma käyttää siinä määriteltyä funktiota `Println`.

```Go
package main
```
Go:ssa jokaisen lähdekooditiedoston alussa on kerrottava minkä nimiseen pakettiin tiedosto kuuluu. Jos paketti on `main`, se on ohjelma, jos nimi on mikä tahansa muu, kyseessä on _kirjasto_.

## Mitä värit tarkoittavat?
