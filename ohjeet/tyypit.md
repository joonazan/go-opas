# Tyypit

Kaikilla arvoilla on tyyppi.

Kuitenkaan kaikki arvot eivät ole samanlaisia. Mitä on `"Hello, world!" + 3`? Entä `3 + fmt.Println`?

Kaikki tyypit rakentuvat yksinkertaisista tyypeistä.

```
bool        sisältää arvon true (tosi) tai false (epätosi)

uint8       etumerkitön  8-bittinen kokonaisluku (0:sta 255:teen)
uint16      etumerkitön 16-bittinen kokonaisluku (0:sta 65535:teen)
uint32      etumerkitön 32-bittinen kokonaisluku (0:sta 4294967295:teen)
uint64      etumerkitön 64-bittinen kokonaisluku (0:sta 18446744073709551615:teen)

int8        etumerkillinen  8-bittinen kokonaisluku (-128:sta 127:teen)
int16       etumerkillinen 16-bittinen kokonaisluku (-32768:sta 32767:teen)
int32       etumerkillinen 32-bittinen kokonaisluku (-2147483648:sta 2147483647:teen)
int64       etumerkillinen 64-bittinen kokonaisluku (-9223372036854775808:sta 9223372036854775807:teen)

float32     IEEE-754 -standardin mukainen 32-bittinen liukuluku
float64     IEEE-754 -standardin mukainen 64-bittinen liukuluku

complex64   kompleksiluku, jonka reaali- ja imaginaariosat ovat typpiä float32
complex128  kompleksiluku, jonka reaali- ja imaginaariosat ovat typpiä float64

byte        toinen nimi uint8:lle
rune        toinen nimi int32:lle; voidaan käyttää yhden merkin varastoimiseen

string      merkkijono, eli tekstiä
```
Tulostamalla asioita `%T`-formatoinnilla, voimme tutkia arvojen tyyppejä. On olemassa melko monimutkaisia tyyppejä:
```Go
package main

import "fmt"

func main() {
	fmt.Printf("%T", fmt.Println)
}
```
tulostaa:
```Go
func(...interface {}) (int, error)
```
