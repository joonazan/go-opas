# Bossien tappaminen

Jokaisen bossintappotehtävän mukana tulee kansio, jossa on yksi testitiedosto (loppuu `_test.go`) ja yksi toinen lähdekooditiedosto. Tehtävänäsi on täydentää siinä tiedostossa oleva funktio niin että se tappaa sille annetun bossin.

$$$cyberdemon/main.go$$$

Kaikilla bosseilla on metodi Kuollut, joka palauttaa `true` jos bossi on kuollut. Siitä voi olla apua. Sitä voi käyttää esim. 

```Go
if demoni.Kuollut() {
    println("demoni on kuollut")
}
```

## Miten tiedän onko algoritmini tarpeeksi väkevä?

Sen, kuoleeko bossi näkee ajamalla testit. Tämä tapahtuu esimerkiksi ylläolevan koodin tapauksessa näin:

```sh
go test github.com/joonazan/go-opas/esimerkit/cyberdemon
```

Aluksi se sanoo:

```sh
--- FAIL: TestTapaDemoni (0.00s)
    tappo_test.go:12: TapaDemoni:n jälkeen demoni oli yhä hengissä!
FAIL
FAIL    github.com/joonazan/go-opas/esimerkit/cyberdemon    0.002s
```

Kun bossi kuolee, tulee:

```sh
ok      github.com/joonazan/go-opas/esimerkit/cyberdemon    0.002s
```
