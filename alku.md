# Alku
Tietokoneet osaavat _ajaa_ yksinkertaisista käskyistä koostuvia ohjelmia. Näitä käskyjä kutsutaan usein _konekieleksi_. Konekielisten käskyjen luetteleminen on kuitenkin työlästä, eikä lopputulos edes aina tee sitä mitä sen piti! Siksi jo 1950-luvulla luotiin ensimmäiset _ohjelmointikielet_.

```
  42cb30:	64 48 8b 0c 25 f8 ff 	mov    rcx,QWORD PTR fs:0xfffffffffffffff8
  42cb37:	ff ff 
  42cb39:	48 8d 44 24 c8       	lea    rax,[rsp-0x38]
  42cb3e:	48 3b 41 10          	cmp    rax,QWORD PTR [rcx+0x10]
  42cb42:	0f 86 77 0c 00 00    	jbe    42d7bf <runtime.chanrecv+0xc8f>
  42cb48:	48 81 ec b8 00 00 00 	sub    rsp,0xb8
  42cb4f:	48 8b 8c 24 c8 00 00 	mov    rcx,QWORD PTR [rsp+0xc8]
  42cb56:	00 
  42cb57:	c6 84 24 e1 00 00 00 	mov    BYTE PTR [rsp+0xe1],0x0
  42cb5e:	00 
  42cb5f:	c6 84 24 e0 00 00 00 	mov    BYTE PTR [rsp+0xe0],0x0
  42cb66:	00 
  42cb67:	31 ed                	xor    ebp,ebp
```
_Konekieltä. Komennot ovat toisessa sarakkeessa. Oikealla olevien sanojen on tarkoitus helpottaa niiden ymmärtämistä._

Nykyään ohjelmia tehdään kirjoittamalla _lähdekoodia_ jollain ohjelmointikielellä. Koodi voidaan muuttaa konekieleksi _kääntäjällä_ tai sitten voidaan käyttää _tulkkia_, joka tekee lähdekoodin kuvaamat asiat. Tähän oppaaseen on valittu kieli nimeltä __Go__, koska se on yksinkertainen, sillä on vaikea tehdä vahingossa muuta kuin mitä halusi tehdä ja sillä ohjelmointi on niin miellyttävää, että monet käyttävät sitä.

## [Miten asennan Go:n koneelleni?](setup.md)
## [Ensimmäinen tehtävä: ensimmäinen ohjelma](helloworld.md)

## Miten edetä tästä
[Etusivulla](README.md) on kaavio, josta voi valita mitä opiskelee seuraavaksi.
