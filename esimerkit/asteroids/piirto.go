package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/joonazan/vec2"
)

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
