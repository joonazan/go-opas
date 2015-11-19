package main

type Tila func() Tila

func Aloita(alku Tila) {

	tila := alku

	for {
		tila = tila()
	}
}
