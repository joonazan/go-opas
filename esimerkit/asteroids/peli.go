package main

import "github.com/joonazan/vec2"

type Peli struct {
	paikat, nopeudet []vec2.Vector

	muodot     [][]vec2.Vector
	muunnokset []vec2.Matrix

	alus Alus
}

func (p *Peli) Päivitä(dt float64, ohjaimet Ohjaimet) {

	for i := range p.nopeudet {
		p.paikat[i] = p.paikat[i].Plus(p.nopeudet[i].Times(dt))
	}

	for i, paikka := range p.paikat {
		p.paikat[i] = wrapVector(paikka)
	}

	p.PäivitäAlus(dt, ohjaimet)
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
