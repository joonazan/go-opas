package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("Mistä haluat laskea neliöjuuren? ")
	var x float64
	fmt.Scan(&x)

	fmt.Println("neliöjuuri", x, "on", math.Sqrt(x))

	fmt.Print("Mitkä kaksi kokonaislukua haluat laskea yhteen? ")
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println("Niiden summa on", a+b)
}
