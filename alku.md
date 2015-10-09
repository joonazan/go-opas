# Alku
Tietokoneet osaavat _ajaa_ yksinkertaisista käskyistä koostuvia ohjelmia. Näitä käskyjä kutsutaan usein _konekieleksi_. Konekielisten käskyjen luetteleminen on kuitenkin työlästä, eikä lopputulos edes aina tee sitä mitä sen piti! Siksi jo 1950-luvulla luotiin ensimmäiset _ohjelmointikielet_.

```
  42cb30:	64 48 8b 0c 25 f8 ff 	mov    rcx,QWORD PTR fs:0xfffffffffffffff8
  42cb37:	ff ff 
  42cb39:	48 8d 44 24 c8       	lea    rax,[rsp-0x38]
  42cb3e:	48 3b 41 10          	cmp    rax,QWORD PTR [rcx+0x10]
  42cb42:	0f 86 77 0c 00 00    	jbe    42d7bf <runtime.chanrecv+0xc8f>
  42cb48:	48 81 ec b8 00 00 00 	sub    rsp,0xb8
  42cb4f:	48 8b 8c 24 c8 00 00 	mov    rcx,QWORD PTR [rsp+0xc8]
  42cb56:	00 
  42cb57:	c6 84 24 e1 00 00 00 	mov    BYTE PTR [rsp+0xe1],0x0
  42cb5e:	00 
  42cb5f:	c6 84 24 e0 00 00 00 	mov    BYTE PTR [rsp+0xe0],0x0
  42cb66:	00 
  42cb67:	31 ed                	xor    ebp,ebp
```
_Konekieltä. Komennot ovat toisessa sarakkeessa. Oikealla olevien sanojen on tarkoitus helpottaa niiden ymmärtämistä._

Nykyään ohjelmia tehdään kirjoittamalla _lähdekoodia_ jollain ohjelmointikielellä. Koodi voidaan muuttaa konekieleksi _kääntäjällä_ tai sitten voidaan käyttää _tulkkia_, joka tekee lähdekoodin kuvaamat asiat. Tähän oppaaseen on valittu kieli nimeltä __Go__, koska se on yksinkertainen, sillä on vaikea tehdä vahingossa muuta kuin mitä halusi tehdä ja sillä ohjelmointi on niin miellyttävää, että monet käyttävät sitä.

## Hello, world!
Ohjelmointi on perinteistä aloittaa ohjelmalla, joka _tulostaa_ `Hello, world!`. Tulostaminen tarkoittaa tekstin ulos antamista, niin että se voidaan näyttää komentorivillä tai tallentaa tiedostoon. Nykyään voi elää näkemättä koskaan komentoriviä(ja ne joita nykyään näkee ovat [virtuaalisia](https://en.wikipedia.org/wiki/Virtual_console)), mutta ohjelmoijalle tulostaminen on yhä helpoin tapa viestiä ohjelmalla ulkomaailmalle ohjelman sisuksista.

Nyt on hyvä vaihe tarkistaa, että pystyt [ajamaan koodia](setup.md). Kokeile ajaa annettu ohjelma.

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

## Miten edetä tästä
Etusivulla on kaavio, josta voi valita mitä opiskelee seuraavaksi.
