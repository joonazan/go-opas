package main

import (
	"github.com/joonazan/vec2"
	"math"
)

func (p *Peli) TeeAlus() {

	p.alus = len(p.paikat)

	p.paikat = append(p.paikat, vec2.Vector{})
	p.nopeudet = append(p.nopeudet, vec2.Vector{})
	p.kulmat = append(p.kulmat, 0)

	kolmio := []vec2.Vector{{0.1, 0}, {-0.1, 0}, {0, 0.2}, {0.1, 0}}

	p.muodot = append(p.muodot, Muoto{
		Id:      p.alus,
		Pisteet: kolmio,
		Väri:    Väri{1, 1, 1},
		Muunnos: vec2.Translation(vec2.Vector{0, -0.06}),
	})

	p.muodot = append(p.muodot, Muoto{
		Id:      p.alus,
		Pisteet: kolmio,
		Väri:    Väri{1, 0.7, 0.3},
		Muunnos: vec2.Translation(vec2.Vector{0, -0.09}).Mul(vec2.Scale(0.4, 0.4).Mul(vec2.Rotation(math.Pi))),
	})
}

func (p *Peli) PäivitäAlus(dt float64, ohjaimet Ohjaimet) {

	const (
		moottorinVahvuus = 0.8
		ratinVahvuus     = 4
	)

	p.kulmat[p.alus] += ohjaimet.Ratti * ratinVahvuus * dt

	if ohjaimet.Kaasu {
		muutos := vec2.Rotation(p.kulmat[p.alus]).Transform(vec2.Vector{0, moottorinVahvuus * dt})

		nopeus := &p.nopeudet[p.alus]
		*nopeus = nopeus.Plus(muutos)
	}
}
