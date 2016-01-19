package main

import "fmt"

func main() {
	summa := 0
	luku := 1
	for luku <= 100 {
		summa += luku
		luku++
	}
	fmt.Println(summa)
}
