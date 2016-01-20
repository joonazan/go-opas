# Funktio

Funktio on tapa antaa nimi jonkin tekemiselle.

Funktion määrittely alkaa aina avainsanalla `func`. Sitten tulee funktion nimi ja sitten sulut. Aaltosulkuihin tulee mitä funktio tekee.

> Eivätkö nimen jälkeen tulevat sulut ole turhat?
> - Tämän sivun asioissa ovat, mutta sivulla Argumentit ja paluuarvot selviää niiden merkitys.

```Go
func melua() {
    fmt.Println("TÖÖT!")
}
```

Tähän mennessä olet tehnyt vain main-nimisiä funktioita. Muiden funktioiden sisään voi kirjoittaa aivan samanlaista koodia. Ainoa ero on, että niitä ei automattisesti kutsuta.

## Apua, laitoin funktioon koodia, mutta se ei tee mitään!

Jotta koodi funktion sisällä suoritettaisiin, funktiota pitää kutsua. Se tehdään kirjoittamalla funktion nimi ja laittamalla sen perään sulut.

Funktioita käyttävä, melko järjetön ohjelma:

$$$func/main.go$$$

Jos haluat nähdä järkevämpää käyttöä funktioille, jatka lukemalla funktioden argumenteista ja paluuarvoista.