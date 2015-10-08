# Funktio

![funktiokone](https://upload.wikimedia.org/wikipedia/commons/thumb/3/3b/Function_machine2.svg/243px-Function_machine2.svg.png)
![kahden joukon funktio](https://upload.wikimedia.org/wikipedia/commons/thumb/d/df/Function_color_example_3.svg/223px-Function_color_example_3.svg.png)

> Funktio on kone jonne laitetaan jotain sisään ja jotain tulee ulos.
>
> – _peruskoulun matematiikan opettaja_

<!-- -->
> Funktio on kuvaus joukosta toiseen.
>
> – _yliopiston matematiikan opettaja_

<!-- -->
> Funktio tekee jotain.
>
> – _minä, ohjelmoinnin opettaja_

Matematiikan funktiota kutsutaan ohjelmoinnissa _puhtaaksi funktioksi_. Funktio ohjelmoinnissa on puhdasta funktiota yksinkertaisempi ja mahtavampi. Mitä `Println` olisi matematiikassa?

Funktion nimi esiintyy kahdessa eri asiayhteydessä: funktiota voi _kutsua_, mutta sitä ennen se pitää olla _määritelty_. Tässä on hieman funktiokutsuja.
```Go
kuutiojuuri(luku)
funktio(argumentti1, argumentti2, argumentti3)
teeJotain()
```
Funktiokutsu koostuu funktion nimestä, jota seuraavat sulut, jotka sisältävät funktion _argumentit_. Argumentit ovat ne arvot, jotka "menevät funktion sisään".

```Go
func teeJotain() {
	teeSe()
	teeKannu()
}
