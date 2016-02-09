package main

import "github.com/joonazan/vec2"

func (p *Peli) TeeAsteroidi() {

	id := len(p.paikat)

	p.nopeudet = append(p.nopeudet, vec2.Vector{1, 0.1})
	p.paikat = append(p.paikat, vec2.Vector{1, 1})
	p.kulmat = append(p.kulmat, 0)
	p.muodot = append(p.muodot, Muoto{
		Id: id,
		Pisteet: []vec2.Vector{
			{0, 0},
			{0.3, 2},
			{-0.4, 2.6},
			{-1.2, 2.3},
			{-2.5, 2.7},
			{-2.8, 2.1},
			{-2, 0.7},
			{-2, 0.4},
			{-1, 0.3},
			{-0.8, 0},
			{0, 0},
		},
		Väri:    Väri{1, 1, 1},
		Muunnos: vec2.Scale(0.1, 0.1).Mul(vec2.Translation(vec2.Vector{1, -1.5})),
	})

	p.pyörimiset = append(p.pyörimiset, Pyöriminen{id, 0.5})
}
