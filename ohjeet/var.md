# Muuttujat

Muuttujaan voi varastoida jotain. Koska eri asiat ovat eri kokoisia, muuttujaan voi varastoida vain yhden [tyyppisiä](tyypit.md) arvoja.

>Mutta esimerkiksi `int8` ja `uint8` -tyyppiset arvothan ovat molemmat 8 bitin pituisia! Miksi niitä ei voi varastoida samaan muuttujaan? Koska silloin voisi tapahtua ikäviä vahinkoja. `0xFF`(kahdeksan ykkösbittiä) on -1 `int8`:na ja 255 `uint8`:na!

Muuttujan voi määritellä laittamalla ensin avainsanan `var`, sitten muuttujan nimen ja lopulta muttujan tyypin. Tällöin muuttujan arvo on kyseisen tyypin [_nolla-arvo_](nollaarvo.md).

```Go
var muuttuja int
```

Yleensä muuttujaan kuitenkin halutaan heti varastoida jotain.

```Go
silmien_määrä := 2
```

Muuttujan arvoa voi käyttää yksinkertaisesti kirjoittamalla sen nimen.

```Go
fmt.Println(silmien_määrä)
```

Muuttujan nimi tulee siitä, että sitä voi muuttaa sijoittamalla siihen uuden arvon.

```Go
silmien_määrä = 3
```
