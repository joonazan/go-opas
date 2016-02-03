package main

import "github.com/joonazan/vec2"

type Peli struct {
	paikat, nopeudet []vec2.Vector

	kulmat []float64
	muodot []Muoto

	pyörimiset []Pyöriminen

	alus int
}

type Pyöriminen struct {
	Id     int
	Nopeus float64
}

func (peli *Peli) Päivitä(dt float64, ohjaimet Ohjaimet) {

	for i := range peli.nopeudet {
		peli.paikat[i] = peli.paikat[i].Plus(peli.nopeudet[i].Times(dt))
	}

	for i, paikka := range peli.paikat {
		peli.paikat[i] = wrapVector(paikka)
	}

	peli.PäivitäAlus(dt, ohjaimet)

	for _, p := range peli.pyörimiset {
		peli.kulmat[p.Id] += p.Nopeus * dt
	}
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
