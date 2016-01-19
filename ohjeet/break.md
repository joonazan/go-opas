# Ikuinen silmukka

Jos forille ei anna mitään ehtoa, se toistaa sisältöään ikuisesti.

```Go
package main

func main() {
    for {
    }
}
```

Tämä on 5/5 ohjelma. Suosittelen lämpimästi.

## Miten silmukka lopetetaan?

`break`-lause menee ulos silmukasta.

$$$break/main.go$$$

Muut tavat päästä ulos silmukasta ovat `return`, `goto` ja paniikki.

## Kokeile itse

Tee ohjelma, joka loppuu kun sille kirjoittaa hei käyttäen breakkia.
Tässä on sellainen kirjoitettuna ilman breakkia:

$$$vaatii_tervehdyksen/main.go$$$