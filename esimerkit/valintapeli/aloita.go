package main

type Tila func() Tila

// Käynnistää pelin annetusta tilasta.
func Aloita(alkutila Tila) {

	tila := alkutila

	for {
		tila = tila()
	}
}
