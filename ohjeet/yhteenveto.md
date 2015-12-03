# Yhteenveto

`{}`
Aaltosulkujen sisällä on koodiblokki. Koodirivit blokissa suoritetaan ylhäältä alas.

`()`
Suluilla kutsutaan funktioita ja ryhmitellään operaattoreita:
```Go
fmt.Println((1+2)*3)
```

`func`
Määrittelee funktion.

`if`
Suorittaa koodiblokin jos ennen blokkia oleva arvo on `true` (tosi).

`for`
Sama kuin if, mutta suorittaa blokkia toistuvasti. Aina ennen blokin suorittamista koodi katsoo, onko keskimmäinen osa tosi.

Jos laittaa ennen koodiblokkia kolme puolipisteellä erotettua osaa, ensimmäinen suoritetaan kun tullaan foriin. Viimeinen suoritetaan jokaisen iteraation jälkeen. Tätä käytetään yleensä laskemiseen.
```Go
for i := 0; i < 20; i++ {fmt.Println(i)}
```
Tämä koodi laskee nollasta 19:ta asti.

`var` tekee uuden muuttujan.
```Go
var muuttuja tyyppi
```
Funktion sisällä on kätevämpää kirjoittaa:
```Go
muuttuja := arvo
```
Muuttujan arvon muuttaminen kun se on luotu onnistuu yksinkertaisesti yhtäsuuruusmerkillä!

Muuttuja näkyy siellä missä se on määritelty, eli ei esimerkiksi toisessa funktiossa, jos se on määritelty yhdessä.
