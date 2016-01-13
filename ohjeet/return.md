# Argumentit ja paluuarvot

Jos funktiosta tulee jotain ulos, sulkujen jälkeen kirjoitetaan _paluuarvon_ [_tyyppi_](tyypit.md). `return`-avainsana lopettaa heti funktion suorituksen. Sen yhteydessä myös annetaan arvot jotka tulevat funktiosta ulos.

```Go
func annaViisi() int {
    return 5
}

func onkoPienempiKuinKolme(luku uint32) bool {
    return luku < 3
}
```

Sulkujen sisään tulevat nimet ja tyypit argumenteille.

```Go
func tervehdi(nimi string) {
    fmt.Println("Hei,", nimi)
}

func tervehdiKohteliaasti(nimi string) {
    tervehdi("arvoisa " + nimi)
}

func summa(a, b int) int {
    return a + b
}
```
