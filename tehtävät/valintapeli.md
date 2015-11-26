# Valintapeli

Tällä kertaa pääset tekemään pelin, jossa valitaan useasta vaihtoehdosta. 

Hanki koodi komennolla `go get github.com/joonazan/go-opas/esimerkit/valintapeli`. Kuten aina, käännä ja aja koodi. Kun olet kokeillut koodia, siirry tehtäviin. Jos tulee tarve ymmärtää lisää siitä miten koodi toimii, lue lisätietoja tehtävien alapuolelta.

## Tehtävät
- Muuta itkemisen teksti joksikin muuksi kuin "yhyy...".
- Lisää yksi vaihtoehto sen, että ei tee mitään, ja itkemisen ohelle.
- Tee kokonaan oma peli. Jos et keksi itse ideaa, tee tarina hirviön päivästä.


## Miten koodi toimii?

Koodi on tällä kertaa monessa tiedostossa. Se toimii juuri samalla tavalla kuin jos kaikki koodi paitsi `package` ja `ìmport` olisi samassa samassa tiedostossa.

Pelin tekemiseen ei tarvitse nähdä kuin main.go. Peli on aina yhdessä tilassa. Tilat ovat [funktioita](../ohjeet/funktio.md), jotka palauttavat seuraavan tilan.

### valitse.go ja aloita.go

Nämä tiedostot sisältävät funktiot `Valitse` ja `Aloita`. Funktion kuten `Valitse` koodaaminen voi näyttää vaikealta, mutta se johtuu siitä, että koodi oli alkuun paljon yksinkertaisempi. Tässä funktion ensimmäinen versio:

```Go
func Valitse(vaihtoehdot []Vaihtoehto) Tila {

    for i, v := range vaihtoehdot {
        fmt.Printf("%c) %s\n", 'a'+i, v.kuvaus)
    }

    var sana string
    for len(sana) == 0 || sana[0] < 'a' || int(sana[0]-'a') >= len(vaihtoehdot) {
        fmt.Scan(&sana)
    }

    return vaihtoehdot[sana[0]-'a'].seuraus
}
```
