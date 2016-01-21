# Muuttujat

Muuttujaan on nimetty pala tietokoneen muistia, jonne voi varastoida arvon.

Seuraava ohjelma laskee hyvin yksinkertaisen laskun: kun tiedetään A0-arkin koko, minkä kokoinen A1-paperiarkki on?

![Kuva paperikoista](https://upload.wikimedia.org/wikipedia/commons/thumb/8/8a/A_size_illustration.svg/444px-A_size_illustration.svg.png)

$$$var/main.go$$$

`:=` luo uuden muuttujan. `=` vaihtaa muuttujan arvoa.

## Miksi `:=`? Eikö sekin voisi olla `=`?

Koska muuten voisi vahingossa kirjoittaa vaikka `levyes` ja ohjelma toimisi väärin ilman mitään näkyvää syytä.

## Mitä varten muuttuja `vanhaLeveys` on?

Se on opettamassa miksi temppu, jonka näytän kohta, on kätevä. Kokeile vaihtaa se kohta koodista tähän:

```Go
leveys = pituus / 2
pituus = leveys
```

Nyt ohjelmasta tulee ulos aivan muuta kuin pitäisi. Mieti hetken miksi. Jos et keksi, keskustele asiasta jonkun kanssa.

```Go
pituus, leveys = leveys, pituus / 2
```

Tämä on se temppu, josta puhuin.

## Kokeile itse
- Kun olet opetellut for-silmukan, muuta ohjelma laskemaan kaikki paperikoot A8 asti.

## Bonus: Tietyntyyppisen muuttujan tekeminen

Jos haluat tehdä muuttujan, joka varastoi tietynlaisia arvoja, mutta et välitä siitä, mitä se aluksi sisältää, voit luoda sen hyvin virallisin elkein:

```Go
var luku int
```

Tuo koodi tekee itse asiassa täsmälleen saman asian kuin tämä:

```Go
luku := 0
```

Tarkemmankin tyypin saa määrättyä käyttäen `:=`-merkintää. Siksi `var nimi tyyppi`-merkintää näkee melko harvoin.

```Go
luku := uint8(0)
```
