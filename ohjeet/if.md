# `if`

```Go
package main

import "fmt"

func main() {
	if true {
		fmt.Println("Tämä näkyy")
	}

	if false {
		fmt.Println("Tämä ei näy")
	}
}

```

Jos `if`:lle annetaan `true`, niin sen sisällä oleva koodi suoritetaan. Jos sille annetaan `false`, niin sitä ei suoriteta. Mitään muuta iffille ei voikaan antaa, koska sille on pakko antaa jotain `bool`-[tyyppistä](tyypit.md).

Tämähän on aivan hyödytöntä! Onneksi iffille voi antaa jotain josta ei tiedä onko se `true` vai `false`. Booleaneja voi rakentaa kaikenlaisista muista arvoista vertailuoperaattoreilla. Niitä on `==`, `!=`, `<`, `>`, `<=` ja `>=`. Ne toimivat juuri niin kuin matematiikassa. `!=` tarkoittaa "ei ole yhtäsuuri".

```Go
if korkeus < 100 * kilometri {
	fmt.Println("Et ole avaruudessa.")
}
```

`!booleani` tarkoittaa käänteistä. Esimerkiksi `!onSyönytMunkin` kertoo, onko joku jättänyt syömättä munkin.

```Go
onAvaruudessa := korkeus >= 100 * kilometri
if onAvaruudessa {
  fmt.Println("Olet avaruudessa.")
}
if !onAvaruudessa {
  fmt.Println("Et ole avaruudessa.")
}
```

Aika hankalaa. Onneksi saman voi kirjoittaa helpommin.

```Go
if korkeus >= 100 * kilometri {
  fmt.Println("Olet avaruudessa.")
} else {
  fmt.Println("Et ole avaruudessa.")
}
```

Samalla tavalla kuin foriin, iffiin voi laittaa koodirivin, joka suoritetaan ennen varsinaista osaa.

```Go
if teksti, virhe := haeSivu(); virhe == nil {
  lueSivu(teksti)
} else {
  fmt.Println("Tuli virhe:", virhe)
}
```
