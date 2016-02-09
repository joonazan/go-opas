package main

import (
	"github.com/joonazan/vec2"
	"math"
)

var alus Alus

type Alus struct {
	Id     int
	Liekki int
}

func TeeAlus() {

	alus.Id = len(paikat)

	paikat = append(paikat, vec2.Vector{})
	nopeudet = append(nopeudet, vec2.Vector{})
	kulmat = append(kulmat, 0)

	kolmio := []vec2.Vector{{0.1, 0}, {-0.1, 0}, {0, 0.2}, {0.1, 0}}

	muodot = append(muodot, Muoto{
		Id:      alus.Id,
		Pisteet: kolmio,
		Väri:    Väri{1, 1, 1},
		Muunnos: vec2.Translation(vec2.Vector{0, -0.06}),
	})

	alus.Liekki = len(muodot)
	liekki := Muoto{
		Id:      alus.Id,
		Pisteet: kolmio,
		Väri:    Väri{1, 0.7, 0.3},
		Muunnos: vec2.Translation(vec2.Vector{0, -0.09}).Mul(vec2.Scale(0.4, 0.4).Mul(vec2.Rotation(math.Pi))),
	}
	muodot = append(muodot, liekki)
}

func PäivitäAlus(dt float64, ohjaimet Ohjaimet) {

	const (
		moottorinVahvuus = 0.8
		ratinVahvuus     = 4
	)

	kulmat[alus.Id] += ohjaimet.Ratti * ratinVahvuus * dt

	if ohjaimet.Kaasu {
		muutos := vec2.Rotation(kulmat[alus.Id]).Transform(vec2.Vector{0, moottorinVahvuus * dt})

		nopeus := &nopeudet[alus.Id]
		*nopeus = nopeus.Plus(muutos)
	}
}
