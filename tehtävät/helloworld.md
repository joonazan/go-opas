## Hello, world!
Ohjelmointi on perinteistä aloittaa ohjelmalla, joka _tulostaa_ `Hello, world!`. Tulostaminen tarkoittaa tekstin ulos antamista, niin että se voidaan näyttää komentorivillä tai tallentaa tiedostoon. Nykyään voi elää näkemättä koskaan komentoriviä(ja ne joita nykyään näkee ovat [virtuaalisia](https://en.wikipedia.org/wiki/Virtual_console)), mutta ohjelmoijalle tulostaminen on yhä helpoin tapa viestiä ohjelmalla ulkomaailmalle ohjelman sisuksista.

Nyt on hyvä vaihe tarkistaa, että [pystyt ajamaan koodia](setup.md). Kokeile ajaa annettu ohjelma.

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
[`Println`](https://golang.org/pkg/fmt/#Println) on [_funktio_](funktio.md), joka saa ohjelman tulostamaan jotain. Se määritellään _paketissa_ [fmt](https://golang.org/pkg/fmt/). Sitä _kutsutaan_ _merkkijonolla_ `Hello, world!`.

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
Koodilaatikoiden ulkopuolella olevat siniset sanat ovat linkkejä; käytä niitä saadaksesi lisää tietoa sanasta.

Laatikoiden sisällä samanlaiset sanat sanat näkyvät samalla värillä. Värityksen tarkoitus on helpottaa koodin lukemista.

- `package`, `import` ja `func` ovat _avainsanoja_, joilla on aina tietty merkitys.
- `"fmt"` ja `"Hello, world!"` ovat _merkkijonoja_.
- koodissa nähty
 - `main` on funktion nimi funktion määrittelyssä
 - `Println` on funktion nimi funktiokutsussa
