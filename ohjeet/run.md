# Koodin ajaminen

Aluksi tarvitset tietenkin koodia mitä ajaa. Koodin, jonka voi ajaa tunnistaa `package main` alkavista tiedostoista. Käytän tässä esimerkkinä kansiota, joka sisältää tämän tiedoston:

$$$aalto/main.go$$$

Koodin tulisi olla go-kansiosi sisällä. Minulla sen polku on `/home/joonazan/go/src/github.com/joonazan/go-opas/esimerkit/aalto/`. Paketin nimi on polku kansion `src` jälkeen. Minulla siis `github.com/joonazan/go-opas/esimerkit/aalto/`.

Paketti voidaan kääntää komennolla `go install`. Minulla siis `go install github.com/joonazan/go-opas/esimerkit/aalto/`. Tämän komennon jälkeen go-kansiosi kansiosta `bin` löytyy binääri. Voit käynnistää sen tuplaklikkaamalla sitä tai kirjoittamalla komentoriviin sen nimen.

## Hirveän hankalaa.

Komentoja ei tarvitse joka kerta kirjoittaa uudestaan; nuolinäppäimillä ylös ja alas saa näkyviin edellisiä komentoja.

Ohjelman kääntämisen ja ajamisen saa kätevästi yhdeksi komennoksi, joka ajaa ohjelman vain jos kääntäminen onnistuu:

```sh
go install github.com/joonazan/go-opas/esimerkit/aalto && aalto
```

## Miten pääsen ajamaan ja muokkaamaan täällä näkyvää koodia?

Saat haettua netistä uusimman version kaikista esimerkkikoodeista komennolla

```sh
go get -d -u github.com/joonazan/go-opas/esimerkit/...
```

Esimerkkikoodit kertovat mistä ne löytyvät. Esimerkiksi tällä sivulla olevan koodin alapuolella lukee _../esimerkit/aalto/main.go_. Paketin nimi on siis `github.com/joonazan/go-opas/esimerkit/aalto/`.

__Aja kaikki esimerkkikoodit joihin törmäät! Vielä parempi jos muokkaat niitä.__
