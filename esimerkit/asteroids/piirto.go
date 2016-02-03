package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/joonazan/vec2"
	"math/rand"
)

type Muoto struct {
	Id int

	Pisteet    []vec2.Vector
	Muunnos    vec2.Matrix
	Väriskeema Väriväli
}

type Väriväli struct {
	min, max [3]float32
}

func (p Peli) Piirrä() {
	for _, muoto := range p.muodot {

		muunnos := vec2.Rotation(p.kulmat[muoto.Id]).Mul(muoto.Muunnos)
		muunnos = vec2.Translation(p.paikat[muoto.Id]).Mul(muunnos)

		gl.Begin(gl.LINE_STRIP)
		for _, piste := range muoto.Pisteet {
			var w [3]float32
			for c := range w {
				x := rand.Float32()
				w[c] = muoto.Väriskeema.min[c]*x + muoto.Väriskeema.max[c]*(1-x)
			}
			gl.Color3f(w[0], w[1], w[2])
			piirräPiste(muunnos.Transform(piste))
		}
		gl.End()
	}
}

func piirräPiste(v vec2.Vector) {
	gl.Vertex2d(v.X, v.Y)
}
