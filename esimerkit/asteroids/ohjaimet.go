package main

import "github.com/go-gl/glfw/v3.1/glfw"

type Ohjaimet struct {

	// 1 on vasemmalle, -1 on oikealle
	Ratti float64

	Kaasu, Liipaisin bool
}

func painalluksetOhjaimiksi(ikkuna *glfw.Window) (ohjaimet Ohjaimet) {

	ohjaimet.Kaasu = ikkuna.GetKey(glfw.KeyUp) == glfw.Press
	ohjaimet.Liipaisin = ikkuna.GetKey(glfw.KeySpace) == glfw.Press

	if ikkuna.GetKey(glfw.KeyLeft) == glfw.Press {
		ohjaimet.Ratti += 1
	}
	if ikkuna.GetKey(glfw.KeyRight) == glfw.Press {
		ohjaimet.Ratti -= 1
	}

	return
}
