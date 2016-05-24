package main

import "github.com/joonazan/vec2"

func TeeAsteroidi() {

	id := TeeEsine(vec2.Vector{0, 1}, vec2.Vector{1, 0.1}, 0)

	mid := muodolle.Varaa(id)
	muodot[mid] = Muoto{
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
	}

	//diskolle.Varaa(mid)

	pyörimiset[pyörimiselle.Varaa(id)] = 0.5
}
