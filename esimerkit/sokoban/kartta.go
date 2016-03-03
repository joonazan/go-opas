package main

import (
	"fmt"
	"os"
)

type Ruutu uint8

const (
	Tyhjä = iota
	Pelaaja
	Seinä
	Kivi
	Kuoppa
	Maali
	KuviaTextuurissa
)

var (
	kartta                      []Ruutu
	kartanLeveys, kartanKorkeus int
	PelaajaX, PelaajaY          int
)

func init() {
	merkkiRuuduksi := map[rune]Ruutu{
		'#': Seinä,
		'.': Tyhjä,
		'o': Kivi,
		'x': Kuoppa,
		'p': Pelaaja,
		'm': Maali,
	}

	file, err := os.Open("testitaso.mp")
	if err != nil {
		panic("En saanut ladattua tasoa: " + err.Error())
	}

	fmt.Fscan(file, &kartanLeveys, &kartanKorkeus, &PelaajaX, &PelaajaY)

	kartta = make([]Ruutu, kartanLeveys*kartanKorkeus)
	for y := 0; y < kartanKorkeus; y++ {
		var rivi string
		fmt.Fscan(file, &rivi)
		for x, merkki := range rivi {
			LaitaKartanKohtaan(x, y, merkkiRuuduksi[merkki])
		}
	}
}

func LueKartanKohta(x, y int) Ruutu {
	return kartta[laskeIndeksi(x, y)]
}

func LaitaKartanKohtaan(x, y int, r Ruutu) {
	kartta[laskeIndeksi(x, y)] = r
}

func laskeIndeksi(x, y int) int {
	if x < 0 || x >= kartanLeveys {
		panic(fmt.Sprintf("x on %d! Kartan leveys on %d.", x, kartanLeveys))
	}
	if y < 0 || y >= kartanKorkeus {
		panic(fmt.Sprintf("y on %d! Kartan korkeus on %d.", y, kartanKorkeus))
	}

	return kartanLeveys*y + x
}
