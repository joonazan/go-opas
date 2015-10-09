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
  42cb69:	48 39 e9             	cmp    rcx,rbp
  42cb6c:	75 6c                	jne    42cbda <runtime.chanrecv+0xaa>
  42cb6e:	80 bc 24 d8 00 00 00 	cmp    BYTE PTR [rsp+0xd8],0x0
  42cb75:	00 
  42cb76:	75 08                	jne    42cb80 <runtime.chanrecv+0x50>
  42cb78:	48 81 c4 b8 00 00 00 	add    rsp,0xb8
  42cb7f:	c3                   	ret    
  42cb80:	48 c7 04 24 00 00 00 	mov    QWORD PTR [rsp],0x0
  42cb87:	00 
  42cb88:	48 c7 44 24 08 00 00 	mov    QWORD PTR [rsp+0x8],0x0
  42cb8f:	00 00 
  42cb91:	48 8d 1d 68 7c 1a 00 	lea    rbx,[rip+0x1a7c68]        # 5d4800 <go.string.*+0x22760>
  42cb98:	48 89 5c 24 10       	mov    QWORD PTR [rsp+0x10],rbx
  42cb9d:	48 c7 44 24 18 17 00 	mov    QWORD PTR [rsp+0x18],0x17
  42cba4:	00 00 
  42cba6:	c6 44 24 20 10       	mov    BYTE PTR [rsp+0x20],0x10
  42cbab:	48 c7 44 24 28 02 00 	mov    QWORD PTR [rsp+0x28],0x2
```

Nykyään ohjelmia tehdään kirjoittamalla _lähdekoodia_ jollain ohjelmointikielellä. Koodi voidaan muuttaa konekieleksi _kääntäjällä_ tai sitten voidaan käyttää _tulkkia_, joka tekee lähdekoodin kuvaamat asiat. Tähän oppaaseen on valittu kieli nimeltä __Go__, koska se on yksinkertainen, sillä on vaikea tehdä vahingossa muuta kuin mitä halusi tehdä ja sillä ohjelmointi on niin miellyttävää, että monet käyttävät sitä.

## Hello, world!
Ohjelmointi on perinteistä aloittaa ohjelmalla, joka _tulostaa_ `Hello, world!`. Tulostaminen tarkoittaa tekstin ulos antamista, niin että se voidaan näyttää komentorivillä tai tallentaa tiedostoon. Nykyään voi elää näkemättä koskaan komentoriviä(ja ne joita nykyään näkee ovat [virtuaalisia](https://en.wikipedia.org/wiki/Virtual_console)), mutta ohjelmoijalle tulostaminen on yhä helpoin tapa viestiä ohjelmalla ulkomaailmalle.

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
