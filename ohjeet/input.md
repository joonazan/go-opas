# Syötteen lukeminen

$$$input/main.go$$$

Aja tämä ohjelma.  
Kirjoita luku ja paina Enteriä.  
Kirjoita kaksi välilyönnillä erotettua lukua.
Paina Enteriä.

Funktio `fmt.Scan` laittaa sille annettuihin muuttujiin käyttäjän antamat välilyönnillä erotetut arvot.

Se, miten käyttäjän antama teksti tulkitaan riippuu muuttujan tyypistä. Esimerkiksi jos muuttujan tyyppi on `string`, `1.4` tulkittaisiin tekstinä "1.4". Muuttujan tyypillä `float64` se tulkittaisiin lukuna 1,4.

Muuttujien edessä on `&`, koska Scan ei tarvitse muuttujien arvoa vaan niiden osoitteen.

## Tiedostojen lukeminen

Käynnistä ohjelma uudestaan.  
Kirjoita kolme välilyönnillä erotettua lukua.
Paina Enteriä.

Ohjelma ottaa mukisematta luvut, jotka annettiin sille ennen kuin se pyysi niitä! Tämä on kätevään, sillä voimme syöttää ohjelmalle tiedoston. Komentorivillä se tehdään näin:

```sh
ohjelman-nimi < tiedoston-polku
```

Kokeile tätä! Saat kirjoitettua tiedoston polun komentoriviin raahaamalla tiedoston komentoriviin.

Jos et jaksa tehdä tiedostoa, voit myös ohjata toisen ohjelman outputin (suomeksi _tuloste_) ohjelman syötteeksi:

```sh
echo "4 2 -2" | ohjelman-nimi
```

Tämä on syy, miksi pystyviivaa kutsutaan putkeksi.

## Entä jos vaan haluaisin tietää mitä kirjoitettiin ennen Enterin painamista?
### Tai mitä tiedostossa luki yhdellä rivillä?

$$$readline/main.go$$$
