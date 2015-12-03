package main

import "fmt"

func teeUlottuvuusvaihtoehto(takaisin Tila) Vaihtoehto {
	return Vaihtoehto{"mene toiseen ulottuvuuteen", func() Tila {
		fmt.Println("Olet toisessa ulottuvuudessa.")
		return Valitse(Vaihtoehdot{
			{"mene takaisin", takaisin},
		})
	}}
}
