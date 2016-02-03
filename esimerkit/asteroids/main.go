package main

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/joonazan/closedgl"
)

func main() {
	ikkuna := closedgl.NewWindow(640, 640, "Asteroids")

	peli := Peli{}
	peli.TeeAlus()

	edellinen_aika := 0.0
	closedgl.RunInWindow(func(aika float64) {

		dt := aika - edellinen_aika
		edellinen_aika = aika

		peli.Päivitä(dt, painalluksetOhjaimiksi(ikkuna))
		peli.Piirrä()
	}, ikkuna)
}

func painalluksetOhjaimiksi(ikkuna *glfw.Window) (ohjaimet Ohjaimet) {

	ohjaimet.Kaasu = ikkuna.GetKey(glfw.KeyUp) == glfw.Press

	if ikkuna.GetKey(glfw.KeyLeft) == glfw.Press {
		ohjaimet.Ratti += 1
	}
	if ikkuna.GetKey(glfw.KeyRight) == glfw.Press {
		ohjaimet.Ratti -= 1
	}

	return
}
