package main

import "github.com/joonazan/closedgl"

func main() {
	ikkuna := closedgl.NewWindow(640, 640, "Asteroids")

	TeeAlus()
	TeeAsteroidi()

	edellinen_aika := 0.0
	closedgl.RunInWindow(func(aika float64) {

		dt := aika - edellinen_aika
		edellinen_aika = aika

		Päivitä(dt, painalluksetOhjaimiksi(ikkuna))
		Piirrä()
	}, ikkuna)
}
