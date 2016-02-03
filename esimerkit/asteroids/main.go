package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/joonazan/closedgl"
	"github.com/joonazan/vec2"
)

func main() {
	ikkuna := closedgl.NewWindow(640, 640, "Asteroids")

	peli := Peli{}

	peli.paikat = append(peli.paikat, vec2.Vector{0, 0})
	peli.nopeudet = append(peli.nopeudet, vec2.Vector{0, 0.1})
	peli.muodot = append(peli.muodot, []vec2.Vector{{0.1, 0}, {-0.1, 0}, {0, 0.2}, {0.1, 0}})

	edellinen_aika := 0.0
	closedgl.RunInWindow(func(aika float64) {

		dt := aika - edellinen_aika
		edellinen_aika = aika

		peli.Päivitä(dt, ikkuna)
		peli.Piirrä()
	}, ikkuna)
}

type Peli struct {
	paikat, nopeudet []vec2.Vector

	muodot [][]vec2.Vector
}

func (p Peli) Päivitä(dt float64, ikkuna *glfw.Window) {

	for i := range p.nopeudet {
		p.paikat[i] = p.paikat[i].Plus(p.nopeudet[i].Times(dt))
	}

	for i, paikka := range p.paikat {
		p.paikat[i] = wrapVector(paikka)
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

func (p Peli) Piirrä() {
	for i, muoto := range p.muodot {
		muunnos := vec2.Translation(p.paikat[i])

		gl.Begin(gl.LINE_STRIP)
		for _, piste := range muoto {
			piirräPiste(muunnos.Transform(piste))
		}
		gl.End()
	}
}

func piirräPiste(v vec2.Vector) {
	gl.Vertex2d(v.X, v.Y)
}
