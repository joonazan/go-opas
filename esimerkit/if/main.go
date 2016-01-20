package main

import "fmt"

func main() {
	fmt.Println("Haluatko voittaa?")
	var vastaus string
	fmt.Scan(&vastaus)
	if vastaus == "kyllä" {
		fmt.Println("Voitit.")
	} else {
		fmt.Println("Voitit silti.")
	}
	fmt.Println("Oletko nyt iloinen?")
}
