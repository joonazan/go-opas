package main

import "fmt"

func main() {
	oppilaanNimi := "Ville"
	oppilaanIkä := 173

	kerroIhmisestä(oppilaanNimi, oppilaanIkä)

	toisenOppilaanNimi := "Kalle"
	toisenOppilaanIkä := oppilaanIkä

	kerroIhmisestä(toisenOppilaanNimi, toisenOppilaanIkä)
}

func kerroIhmisestä(nimi string, ikä int) {
	fmt.Printf("%s on %d sekuntia vanha.\n", nimi, ikä)
}
