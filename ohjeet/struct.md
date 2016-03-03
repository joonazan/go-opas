# Struct

Usein on kätevää olla yksi muuttuja monen sijaan. Structilla saa ryhmiteltyä monta arvoa yhteen muuttujaan.

??rinnan
$$$ilman_structia/main.go$$$
??v
$$$structin_kanssa/main.go$$$
??l

Samanlainen ohjelma ilman structeja ja niiden kanssa rinnakkain. Structi toimii aivan kuin sen jokainen osa olisi muuttuja.

??rinnan
```Go
type Ihminen struct{
    Nimi string
    Ikä int
}
```
??v
Määritellään oma tyyppi `Ihminen`, joka on structi, jossa on nimi, joka on merkkijono ja ikä, joka on kokonaisluku.
??l

??rinnan
```Go
oppilas := Ihminen{}
oppilas.Nimi = "Ville"
oppilas.Ikä = 173
```
??v
```Go
oppilas := Ihminen{
    Nimi: "Ville",
    Ikä: 173,
}
```
??v
```Go
oppilas := Ihminen{"Ville", 173}
```
??l

Kolme eri tapaa luoda uusi ihminen, jonka nimi on Ville ja ikä on 173.
