package main

import "github.com/joonazan/vec2"

var (
	paikat, nopeudet []vec2.Vector
	kulmat           []float64

	pyörimiset []Pyöriminen
)

type Pyöriminen struct {
	Id     int
	Nopeus float64
}

func Päivitä(dt float64, ohjaimet Ohjaimet) {

	for i := range nopeudet {
		paikat[i] = paikat[i].Plus(nopeudet[i].Times(dt))
	}

	for i, paikka := range paikat {
		paikat[i] = wrapVector(paikka)
	}

	PäivitäAlus(dt, ohjaimet)

	for _, p := range pyörimiset {
		kulmat[p.Id] += p.Nopeus * dt
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
