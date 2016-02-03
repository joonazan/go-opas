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

	peli.paikat = append(peli.paikat, vec2.Vector{})
	peli.nopeudet = append(peli.nopeudet, vec2.Vector{})
	peli.muodot = append(peli.muodot, []vec2.Vector{{0.1, 0}, {-0.1, 0}, {0, 0.2}, {0.1, 0}})
	peli.muunnokset = append(peli.muunnokset, vec2.Matrix{})

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

	muodot     [][]vec2.Vector
	muunnokset []vec2.Matrix

	aluksenKulma float64
}

func (p *Peli) Päivitä(dt float64, ikkuna *glfw.Window) {

	for i := range p.nopeudet {
		p.paikat[i] = p.paikat[i].Plus(p.nopeudet[i].Times(dt))
	}

	for i, paikka := range p.paikat {
		p.paikat[i] = wrapVector(paikka)
	}

	ratti := 0.0

	if ikkuna.GetKey(glfw.KeyLeft) == glfw.Press {
		ratti += 1
	}
	if ikkuna.GetKey(glfw.KeyRight) == glfw.Press {
		ratti -= 1
	}

	p.aluksenKulma += ratti * 0.1

	var moottorinVahvuus = 0.01

	if ikkuna.GetKey(glfw.KeyUp) == glfw.Press {
		muutos := vec2.Rotation(p.aluksenKulma).Transform(vec2.Vector{0, moottorinVahvuus})

		nopeus := &p.nopeudet[0]
		*nopeus = nopeus.Plus(muutos)
	}

	p.muunnokset[0] = vec2.Rotation(p.aluksenKulma).Mul(vec2.Translation(vec2.Vector{0, -0.06}))
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
		muunnos := vec2.Translation(p.paikat[i]).Mul(p.muunnokset[i])

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
