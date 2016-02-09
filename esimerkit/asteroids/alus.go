package main

import (
	"github.com/joonazan/vec2"
	"math"
)

var alus Alus

type Alus struct {
	ID       int
	LiekinID int
}

var kolmionPisteet = []vec2.Vector{{0.1, 0}, {-0.1, 0}, {0, 0.2}, {0.1, 0}}

var Liekki = Muoto{
	Pisteet: kolmionPisteet,
	Väri:    Väri{1, 0.7, 0.3},
	Muunnos: vec2.Translation(vec2.Vector{0, -0.09}).Mul(vec2.Scale(0.4, 0.4).Mul(vec2.Rotation(math.Pi))),
}

func TeeAlus() {
	alus.ID = TeeEsine(vec2.Vector{}, vec2.Vector{}, 0)

	id := muodolle.Varaa(alus.ID)
	muodot[id] = Muoto{
		Pisteet: kolmionPisteet,
		Väri:    Väri{1, 1, 1},
		Muunnos: vec2.Translation(vec2.Vector{0, -0.06}),
	}

	alus.LiekinID = muodolle.Varaa(alus.ID)
}

func PäivitäAlus(dt float64, ohjaimet Ohjaimet) {

	const (
		moottorinVahvuus = 0.8
		ratinVahvuus     = 4
	)

	kulmat[alus.ID] += ohjaimet.Ratti * ratinVahvuus * dt

	if ohjaimet.Kaasu {
		muutos := vec2.Rotation(kulmat[alus.ID]).Transform(vec2.Vector{0, moottorinVahvuus * dt})

		nopeus := &nopeudet[alus.ID]
		*nopeus = nopeus.Plus(muutos)

		muodot[alus.LiekinID] = Liekki
	} else {
		muodot[alus.LiekinID] = Muoto{}
	}

	if ohjaimet.Liipaisin {
		TeeLuoti(alus.ID)
	}
}

func TeeLuoti(alusID int) {
	aluksenRotaatio := vec2.Rotation(kulmat[alusID])
	keulanSuunta := aluksenRotaatio.Transform(vec2.Vector{0, 1})
	id := TeeEsine(paikat[alusID], keulanSuunta.Plus(nopeudet[alusID]), 0)

	mid := muodolle.Varaa(id)
	muodot[mid] = Muoto{
		Pisteet: []vec2.Vector{{0, 0}, {0, 0.01}},
		Väri:    Väri{1, 1, 1},
		Muunnos: aluksenRotaatio,
	}

	eliniät[eliniälle.Varaa(id)] = 60
}
