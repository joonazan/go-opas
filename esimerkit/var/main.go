package main

import "fmt"

func main() {

	// A0-paperiarkin koko
	pituus := 1189
	leveys := 841

	// Kun arkin pituus puolitetaan, siitä tulee leveys
	vanhaLeveys := leveys
	leveys = pituus / 2
	pituus = vanhaLeveys

	fmt.Println("A1-arkin pinta-ala on:", pituus*leveys)
	fmt.Println("Sen pituus on:", pituus)
	fmt.Println("Sen leveys on:", leveys)
}
