package main

import "github.com/joonazan/vec2"

type Alus struct {
	Kulma float64
	Id    int
}

func (p *Peli) TeeAlus() {

	p.alus.Id = len(p.paikat)

	p.paikat = append(p.paikat, vec2.Vector{})
	p.nopeudet = append(p.nopeudet, vec2.Vector{})
	p.muodot = append(p.muodot, []vec2.Vector{{0.1, 0}, {-0.1, 0}, {0, 0.2}, {0.1, 0}})
	p.muunnokset = append(p.muunnokset, vec2.Matrix{})
}

func (p *Peli) PäivitäAlus(dt float64, ohjaimet Ohjaimet) {

	const (
		moottorinVahvuus = 0.8
		ratinVahvuus     = 2
	)

	p.alus.Kulma += ohjaimet.Ratti * ratinVahvuus * dt

	if ohjaimet.Kaasu {
		muutos := vec2.Rotation(p.alus.Kulma).Transform(vec2.Vector{0, moottorinVahvuus * dt})

		nopeus := &p.nopeudet[p.alus.Id]
		*nopeus = nopeus.Plus(muutos)
	}

	p.muunnokset[p.alus.Id] = vec2.Rotation(p.alus.Kulma).Mul(vec2.Translation(vec2.Vector{0, -0.06}))
}
