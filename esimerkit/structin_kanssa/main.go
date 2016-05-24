package main

import "fmt"

type Ihminen struct {
	Nimi string
	Ikä  int
}

func main() {
	oppilas := Ihminen{"Ville", 173}
	kerroIhmisestä(oppilas)

	toinenOppilas := Ihminen{"Kalle", oppilas.Ikä}
	kerroIhmisestä(toinenOppilas)
}

func kerroIhmisestä(ihminen Ihminen) {
	fmt.Printf("%s on %d sekuntia vanha.\n", ihminen.Nimi, ihminen.Ikä)
}
