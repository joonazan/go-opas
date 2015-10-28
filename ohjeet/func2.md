# Edistynyt funktioiden määritteleminen

Laittamalla tyypin eteen ... voi tehdä funktion, joka ottaa vaikka kuinka monta argumenttia sliceen. Tällaista funktiota kutsutaan _variadiseksi_.

```Go
func kerroArgumenteista(argumentit ...interface{}) {
    fmt.Println("Tälläista minulle annettiin:")
    fmt.Println(argumentit)
}
```

Funktion paluuarvoille voi antaa nimet, jolloin niitä ei tarvitse antaa returnissa. ks. funktio `Jaa`.

```Go
package main

import "fmt"

func nollallajako() string {
    return "Yritit jakaa nollalla!"
}

var Nollallajako = Virhefunktio(nollallajako)

type Virhefunktio func() string

func (v Virhefunktio) Error() string {
    return v()
}

func Jaa(a, b int) (tulos int, virhe error) {
    if b == 0 {
        virhe = Nollallajako
    } else {
        tulos = a / b
    }
    return
}

func main() {
    tulostaArvoTaiVirhe(Jaa(1, 0))
    tulostaArvoTaiVirhe(Jaa(4, 2))
}

func tulostaArvoTaiVirhe(arvo interface{}, virhe error) {
    if virhe != nil {
        fmt.Println(virhe)
    } else {
        fmt.Println(arvo)
    }
}
```
