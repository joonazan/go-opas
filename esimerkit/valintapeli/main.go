package main

import "fmt"

func main() {
	Aloita(alku)
}

func alku() Tila {
	fmt.Println("Sinulla on tylsää. Mitä teet?")
	return Valitse([]Vaihtoehto{
		{"älä tee mitään", alku},
		{"itke", yhyy},
		teeUlottuvuusvaihtoehto(alku),
	})
}

func yhyy() Tila {
	fmt.Println("yhyy...")
	return alku
}
