# Muuttujat

Muuttujaan voi varastoida jotain. Koska eri asiat ovat eri kokoisia, muuttujaan voi varastoida vain yhden [tyyppisiä](tyypit.md) arvoja.

>Mutta esimerkiksi `int8` ja `uint8` -tyyppiset arvothan ovat molemmat 8 bitin pituisia! Miksi niitä ei voi varastoida samaan muuttujaan? Koska silloin voisi tapahtua ikäviä vahinkoja. `0xFF`(kahdeksan ykkösbittiä) on -1 `int8`:na ja 255 `uint8`:na!

Muuttujan voi määritellä laittamalla ensin avainsanan `var`, sitten muuttujan nimen ja lopulta muttujan tyypin. Tällöin muuttujan arvo on kyseisen tyypin _nollaarvo_.

```Go
var muuttuja tyyppi
```

tyyppi | nollaarvo
-------|-------------
luvut | `0`
`string` | `""`
slice tai `map` | `nil`
`interface` | `nil`
`func` | `nil`

```Go

```