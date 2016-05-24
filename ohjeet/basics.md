# Kaikki mitä tarvitsee tietää ohjelmoidakseen

Kaikki muu on vain hienostelua.

Toisaalta voisi sanoa, että tämä on se mitä tarvitsee tietää ollakseen vaarallinen. Jos osaa nämä asiat, osaa tehdä kaiken, mutta ei osaa tehdä mitään kauniisti.

## Tietorakenteet

### Muuttuja

??rinnan
```Go
var rahaaTilillä int
rahaaTilillä = 500
```
??v
```Go
var rahaaTilillä = 500
```
??v
```Go
rahaaTilillä := 500
```
??l

Kaikki kolme luovat uuden kokonaislukuja varastoivan muuttujan `rahaaTilillä` ja tallentavat siihen luvun 500.

??rinnan
```Go
fmt.Println("Tilillä on", rahaaTilillä, "euroa.")
```
??v
Muuttujan nimen laittaminen jonnekin toimii aivan kuin siihen kohtaan olisi laitettu muuttujan sisältämä arvo.
??l

??rinnan
```Go
rahaaTilillä = rahaaTilillä - 34
```
??v
Tietenkin muuttujan arvoa voi muuttaa. (Esimerkiksi jos tililtä pitää maksaa laskuja.)
??l

### Slice

??rinnan
```Go
sanoja := []string{"hei", "kissa", "kara"}
```
??v
```Go
sanoja := make([]string, 3)
sanoja[0] = "hei"
sanoja[1] = "kissa"
sanoja[2] = "kara"
```
??l

Kaksi eri tapaa luoda tietynpituinen slice ja laittaa sinne sanoja.

## Control Flow

??rinnan
```Go
if ehto {
    koodiblokki
}
```
??v
Jos ehto, niin koodiblokki.
??l

??rinnan
```Go
for {
    koodiblokki
}
```
??v
Toista koodiblokkia ikuisesti.
??l

??rinnan
```Go
for ehto {
    koodiblokki
}
```
??v
Toista koodiblokkia niin kauan kuin ehto.
??l

## Funktio

??rinnan
```Go
func funktionNimi(){
    koodiblokki
}
```
*funktion määrittely*
??v
```Go
funktionNimi()
```
*funktiokutsu*
??v
Jos funktio on määritelty jossain kohdassa koodiasi, sen kutsuminen suorittaa määrittelyssä annetun koodiblokin.
??l


# Ylimääräistä?
## Koodiblokki

??rinnan
```Go
a
b
c
```
??v
Koodiblokissa on yksinkertaisesti peräkkäisiä koodirivejä. Ensin tapahtuu a, sitten b ja lopulta c.
??l

??rinnan
```Go
for alustus; ehto; muutos {
    koodiblokki
}
```
??v
Ensin alustus, sitten toista koodiblokkia niin kauan kuin ehto, tehden joka toiston jälkeen muutos.
??l