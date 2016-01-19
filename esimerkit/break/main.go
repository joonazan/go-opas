package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Ilman tätä noppa käyttäytyy joka kerta samalla tavalla
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Println("Nopasta tuli", rand.Intn(6)+1)

		fmt.Print("Haluatko lisää nopanheittoja? [k/e]")
		var kirjain rune
		fmt.Scan(&kirjain)
		if kirjain == 'e' {
			break
		}
	}
}
