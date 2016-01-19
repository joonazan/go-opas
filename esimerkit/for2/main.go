package main

import "fmt"

func main() {
	summa := 0
	for luku := 1; luku <= 100; luku++ {
		summa += luku
	}
	fmt.Println(summa)
}
