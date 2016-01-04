# Ohjelman kulku

Tietokoneohjelma antaa tarkat ohjeet, jotka kertovat tietokoneelle mitä tehdä.

Keittokirjassa on ohje kakulle, jonka avulla kuka tahansa voi leipoa sen. 

Jos vaiheet tekee väärässä järjestyksessä ei kakusta tule kovin maukasta. Esimerkiksi, jos vuuan laittaa uuniin ennen kuin kakkutaikina on siellä, kakku jää raa-aksi. Siksi ohjeita seurataan ylhäältä alas.

Mitä tekisit, jos sinun käskettäisiin leikata kakku kahtia, mutta sinulla ei vielä ole edes taikinaa? Go sanoo silloin näin: `prog.go:6: undefined: kakku`.

Entä jos reseptissä käsketään hakea mansikka, mutta sillä ei tehdä mitään? Reseptin kirjoittaja on selvästi unohtanut jotain. Tästä syystä Go sanookin `prog.go:5: mansikka declared and not used` eikä käännä ohjelmaa.

Mitä ihmettä pitäisi tehdä, kun kakun ohjeessa käsketään tehdä täyte? Mikä ihmeen täyte? Toisesta kohdasta reseptiä löytyykin onneksi erillinen ohje täytteelle. Ohjelmassa "tee täyte" olisi funktiokutsu ja täytteen resepti olisi funktio.

Jotta resepti ei olisi turhan pitkä, se ei kerro miten munan valkuainen ja keltuainen erotetaan toisistaan. Tämän voi kuitenkin lukea kirjan takaa. Go:ssa `import` kertoo, että ohjelma käyttää toisessa paketissa määriteltyjä nimiä.

Kirjan takanakaan ei voida kertoa kaikkea, joten kirja viittaa kokkaamisen alkeisteoksiin. Go:ssa voi käyttää standardikirjaston ulkopuolista pakettia kertomalla mistä sen löytää. Esim. `import "github.com/joonazan/closedgl"`

