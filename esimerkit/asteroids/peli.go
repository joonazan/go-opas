package main

import (
	"github.com/joonazan/vec2"
)

var (
	paikat, nopeudet []vec2.Vector
	kulmat           []float64

	kierrättäjä = UusiKierrättäjä(&paikat, &nopeudet, &kulmat)

	kuolleetIDt []int

	pyörimiset   []float64
	pyörimiselle = UusiKomponentti(&pyörimiset)

	eliniät   []int
	eliniälle = UusiKomponentti(&eliniät)
)

func Tapa(id int) {
	kuolleetIDt = append(kuolleetIDt, id)
}

type Pyöriminen struct {
	Id     int
	Nopeus float64
}

func TeeEsine(paikka, nopeus vec2.Vector, kulma float64) (id int) {

	id = kierrättäjä.Varaa()

	paikat[id] = paikka
	nopeudet[id] = nopeus
	kulmat[id] = kulma
	return
}

func Päivitä(dt float64, ohjaimet Ohjaimet) {

	for i := range nopeudet {
		paikat[i] = paikat[i].Plus(nopeudet[i].Times(dt))
	}

	for i, paikka := range paikat {
		paikat[i] = wrapVector(paikka)
	}

	PäivitäAlus(dt, ohjaimet)

	for i, p := range pyörimiset {
		kulmat[pyörimiselle.Huoltajat[i]] += p * dt
	}

	for i, e := range eliniät {
		if e == 1 {
			Tapa(eliniälle.Huoltajat[i])
		}
		eliniät[i]--
	}

	muodolle.PoistaKuolleet(kuolleetIDt)
	pyörimiselle.PoistaKuolleet(kuolleetIDt)
	eliniälle.PoistaKuolleet(kuolleetIDt)

	for _, id := range kuolleetIDt {
		kierrättäjä.Vapauta(id)
	}
	kuolleetIDt = kuolleetIDt[:0]
}

func wrapVector(v vec2.Vector) vec2.Vector {
	return vec2.Vector{wrap(v.X), wrap(v.Y)}
}

func wrap(x float64) float64 {
	if x > 1 {
		return x - 2
	}
	if x < -1 {
		return x + 2
	}
	return x
}
