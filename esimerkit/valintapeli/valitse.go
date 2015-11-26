package main

import (
	"errors"
	"fmt"
)

type Vaihtoehto struct {
	kuvaus  string
	seuraus Tila
}

func Valitse(v Vaihtoehdot) Tila {

	v.Näytä()

	for {
		var sana string
		fmt.Scanln(&sana)
		tila, err := v.Valitse(sana)

		if err == nil {
			return tila
		}
		fmt.Println(err)
	}
}

type Vaihtoehdot []Vaihtoehto

func (v Vaihtoehdot) Näytä() {
	for i, e := range v {
		fmt.Printf("%c) %s\n", 'a'+i, e.kuvaus)
	}
}

func (v Vaihtoehdot) Valitse(kirjoitettu string) (tila Tila, err error) {

	if len(kirjoitettu) != 1 {
		err = errors.New("Kirjoita yksi kirjain.")
		return
	}

	kirjain := kirjoitettu[0]
	valittu := kirjain - 'a'

	if kirjain < 'a' || int(valittu) >= len(v) {
		err = errors.New(fmt.Sprintf("%c ei ole välillä %c-%c.", kirjain, 'a', 'a'+len(v)-1))
		return
	}

	tila = v[valittu].seuraus
	return
}
