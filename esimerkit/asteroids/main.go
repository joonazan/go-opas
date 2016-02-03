package main

import "github.com/joonazan/closedgl"

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
