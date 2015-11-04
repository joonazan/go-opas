# Salasana

Tällä kertaa pääset tekemään ohjelman ihan alusta lähtien. Tarkoituksena on tehdä ohjelma, joka kysyy salasanaa kunnes se menee oikein. Jos salasana menee oikein, niin kerrotaan salaista tietoa.

Tältä ohjelma näyttää minun koneella ajettuna. Olen kirjoittanut apinapallo ja painanut enteriä jne.
```
[joonazan@arkkikaari ~]$ salasana
anna salasana: apinapallo
väärä salasana
anna salasana: korvapuusti
väärä salasana
anna salasana: oikeinhevonenpatteriniitti
väärä salasana
anna salasana: salasana
NTN:n päiväkirja, 12.12.2012:
Piilotin munkin silinterihatun alle.
[joonazan@arkkikaari ~]$
```

Jos osaat tehdä tällaisen ohjelman, niin siitä vaan. Jos et, niin jatka lukemista. Mutta tee ensin niin pitkälle kuin osaat.

### Lisäominaisuuksia, tee kun olet saanut kaiken toimimaan.

- Laita ohjelma valittamaan, jos tulee liian monta väärää arvausta.
- Tällä hetkellä salasanan voi selvittää katsomalla koodia. Voit kuitenkin laittaa koodiin salasanan sijaan sen tiivisteen, ja verrata sitä kirjoitetun tekstin tiivisteeseen, sillä samalle tekstille tiiviste on aina sama. Tiiviste on englanniksi hash tai checksum. Saat tekstin adler32-tiivisteen `adler32.Checksum([]byte(teksti))`. Kirjaston `adler32` saa `import "hash/adler32"`

## Ohjelman tekeminen

Luo uusi kansio nimeltä `salasana` kansion `go/src` alle. Luo sinne tiedosto `main.go`. Voit toki nimetä tiedoston tai kansion jollain muullakin tavalla.

Kirjoita tiedostoon:
```Go
package main

func main() {

}
```

Saat käännettyä ohjelman komentoriviltä komennolla `go install salasana` ja ajettua sen komennolla `salasana`. Kokeile! Sen ei pitäisi vielä tehdä mitään.

## Salasanan kysyminen

`main`-[funktiota](../ohjeet/funktio.md) kutsutaan kun ohjelma käynnistyy, joten aletaan laittamaan koodirivejä sen sisään, jotta jotain tapahtuisi.

Laita `main`in sisään `fmt.Print("tekstiä")`. Se tulostaa tekstiä ilman että tulee uutta riviä. Jotta sitä voidaan käyttää, tarvitaan `fmt`. Haetaan se siis laittamalla `package main`:n alle `import "fmt"`.

Seuraavat rivit luovat [muuttujan](../ohjeet/var.md) `kirjoitettu` ja tallentavat siihen kirjoitetun sanan.

```Go
var kirjoitettu string
fmt.Scan(&kirjoitettu)
```

## Salasanan tarkistaminen

Ensin pitäisi katsoa, onko se mitä kirjoitettiin salasana. Se onnistuu operaattorilla `==`. Kokeile lisätä koodirivi `fmt.Println(kirjoitettu == "salasana")` ja ajaa ohjelma!

Mutta oikeastihan haluttiin kertoa salaista tietoa, eikä `true`, jos salasana menee oikein. Tähän tarvitset [if-lauseen](../ohjeet/if.md).

## Kysymyksen toistaminen

Nyt ohjelma ei kysy uudestaan, vaikka salasana menisikin väärin. Tämän voi korjata [for-silmukalla](../ohjeet/for.md).

Mutta kun menee oikein, kyseleminen pitäisi lopettaa! Tämä hoituu laittamalla yhdelle riville `break`. Se tarkoittaa "mene ulos for-silmukasta".
