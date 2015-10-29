# Kirjastot

Kirjastoista voi lainata kirjoja. Tai no nykyään sieltä voi lainata vaikka soutuveneen. Mutta keskitytään nyt vanhoihin kunnon aikoihin, kun sieltä sai vain sivistyneitä kirjoja (ja sarjakuvia). Hyvinvarustelluista kirjastoista saa funktioita, tyyppejä, interfaceja ja metodeja.

Niitä voi lainata sieltä, mutta on ensin mentävä sinne kirjastoon. Siksi pitää ensin kirjoittaa import ja ne kirjastot. Sitten tietokone ratkaisee kauppamatkustajan ongelman ja laskee optimaalisen kirjastoreitin. Kun olet kertonut missä kirjastoissa käyt, pitää vielä suunnitella mitä kirjoja sieltä lainaa, koska eihän sitä turhaan kirjastoon halua raahautua.
```Go
package main

import (
	"fmt"
	"kirjasto"
)

func main() {
	fmt.Println(kirjasto.Kirja())
}
```
Kannattaa myös kertoa, mistä kirjastosta niitä kirjoja lainaa, koska samannimisiä kirjoja voi löytyä eri kirjastoista. Jos esimerkiksi lainaat kirjan vahingossa keittokirjastosta, vaikka se oli tarkoitus lainata pornokirjastosta, kokemus ei varmastikaan ole nautinnollinen.

Kirjastoja voi myös perustaa itse, jos on aikaa ja tilaa. Jotta kukaan voisi käyttää sitä sen on oltava kuitenkin ~~Githubissa~~ HelMetissä. Tässä on minun OMA MAHTAVA kirjastoni. Kirjoitin sen IHAN ITSE!
```Go
package kirjasto

func Kirja() int {
	return 42
}
```
Kukaan muu kuin minä itse ei kuitenkaan koskaan arvostaisi sitä tarpeeksi, se on ainoastaan tietokoneellani src-kansiossa omassa kansiossaan. Minulla on siellä myös toinen paljon hienompi kirjasto siellä, mutta en näytä sitä teille.
