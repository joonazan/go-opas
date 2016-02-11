package main

import "math/rand"

var (
	diskolle = maailma.LuoKomponentti(&muodolle.Kuolleet)
)

func PäivitäDisko() {
	for _, v := range diskolle.Vanhemmat {
		muodot[v].Väri = [3]float32{rand.Float32(), rand.Float32(), rand.Float32()}
	}
}
