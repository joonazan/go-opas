# Go:n asentaminen

## Mitä kaikkea pitää tehdä?

`go` on ohjelma, joka muuttaa Go-koodin tietokoneelle sopivaksi, eli Go-kääntäjä.

Sen lisäksi tarvitaan `git`. Go käyttää sitä koodin automaattisesti netistä hakemiseen `go get` -komennolla. Vaikka oma koodisi ei olisi netissä, se todennäköisesti käyttää muiden kirjoittamia paketteja, jotka saa netistä.

`GOPATH` kertoo kansion, josta `go` etsii paketteja. `GOPATH/bin/` sisältää `go install`-komennon tuottamia ajettavia tiedostoja. Jos lisäät `GOPATH/bin`:n `PATH`-ympäristömuuttujaasi, voit ajaa ohjelmiasi komentorivistä kirjoittamalla niiden nimen.

Monet Go-paketit(mm. ne mitä tässä oppaassa käytetään ikkunan avaamiseen ja asioiden piirtämiseen) sisältävät C-kielellä kirjoitettuja osia. Niiden käyttäminen vaatii C-kääntäjän `gcc`.

Kokeile asennusta jollain ohjelmalla, esimerkiksi jollain kurssin esimerkkiohjelmista.

## Windows

Lataa `go`, `git` ja TDM-GCC allaolevista linkeistä.

https://golang.org/dl/  
https://git-scm.com/downloads  
http://tdm-gcc.tdragon.net/  

Suorita Go:n ja TDM-GCC:n asennusohjelmat oletusasetuksilla. Valitse Gitin asennusohjelmassa että haluat käyttää Gittiä komentoriviltä(Windows CMD), muuten oletusasetukset.

Säädä ympäristömuuttujat graafisella työkalulla tai aja järjestelmävalvojakomentorivilla seuraavat komennot:
```
setx GOPATH C:\\Users\Public\go
setx PATH "C:\\Users\Public\go\bin"
```
Voit tietenkin käyttää jotain muuta kansiota kuin `C:\\Users\Public\go`, kunhan kansiota pystyy muokkaamaan ilman järjestelmävalvojaoikeuksia.

## Linux, ohjeet tehty Arch Linuxille

Asenna git ja go.

```sh
sudo pacman -S git go
```

Lisää `.bash_profile` tai `.zshrc`-tiedostosi loppuun:
```
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin
```

Jos `go` valittaa `gcc`:n puutteesta, asenna `base-devel` tai `gcc`.
