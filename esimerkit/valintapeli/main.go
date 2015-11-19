package main

import "fmt"

func main() {
	Aloita(alku)
}

func alku() Tila {
	fmt.Println("Sinulla on tylsää. Mitä teet?")
	return valitse([]Vaihtoehto{
		{"älä tee mitään", alku},
		{"itke", yhyy},
	})
}

func yhyy() Tila {
	fmt.Println("yhyy...")
	return alku
}
