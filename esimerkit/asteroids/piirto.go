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

		verteksit := make([]verteksi, len(muoto.Pisteet))

		for i, piste := range muoto.Pisteet {
			for c := 0; c < 3; c++ {
				x := rand.Float32()
				verteksit[i].väri[c] = muoto.Väriskeema.min[c]*x + muoto.Väriskeema.max[c]*(1-x)
			}
			verteksit[i].piste = muunnos.Transform(piste)
		}

		for _, siirto := range []vec2.Vector{{0, 0}, {0, 2}, {0, -2}, {2, 0}, {-2, 0}} {
			piirräVerteksitSiirrolla(verteksit, siirto)
		}
	}
}

func piirräVerteksitSiirrolla(verteksit []verteksi, siirto vec2.Vector) {
	gl.Begin(gl.LINE_STRIP)
	for _, v := range verteksit {
		v.piste.Add(siirto)
		v.Piirrä()
	}
	gl.End()
}

type verteksi struct {
	piste vec2.Vector
	väri  [3]float32
}

func (v verteksi) Piirrä() {
	gl.Color3f(v.väri[0], v.väri[1], v.väri[2])
	piirräPiste(v.piste)
}

func piirräPiste(v vec2.Vector) {
	gl.Vertex2d(v.X, v.Y)
}
