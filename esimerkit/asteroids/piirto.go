package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/joonazan/vec2"
)

var (
	muodot []Muoto

	muodolle = maailma.LuoKomponentti(&kuolleetIDt, &muodot)
)

type Muoto struct {
	Pisteet []vec2.Vector
	Muunnos vec2.Matrix
	Väri    Väri
}
type Väri [3]float32

func Piirrä() {
	for i, muoto := range muodot {

		muunnos := vec2.Rotation(kulmat[muodolle.Vanhemmat[i]]).Mul(muoto.Muunnos)
		muunnos = vec2.Translation(paikat[muodolle.Vanhemmat[i]]).Mul(muunnos)

		muunnetutPisteet := make([]vec2.Vector, len(muoto.Pisteet))
		for j, piste := range muoto.Pisteet {
			muunnetutPisteet[j] = muunnos.Transform(piste)
		}

		asetaVäri(muoto.Väri)
		for _, siirto := range []vec2.Vector{{0, 0}, {0, 2}, {0, -2}, {2, 0}, {-2, 0}} {
			piirräPisteetSiirrolla(muunnetutPisteet, siirto)
		}
	}
}

func piirräPisteetSiirrolla(pisteet []vec2.Vector, siirto vec2.Vector) {
	gl.Begin(gl.LINE_STRIP)
	for _, p := range pisteet {
		piirräPiste(p.Plus(siirto))
	}
	gl.End()
}

func piirräPiste(v vec2.Vector) {
	gl.Vertex2d(v.X, v.Y)
}

func asetaVäri(väri Väri) {
	gl.Color3f(väri[0], väri[1], väri[2])
}
