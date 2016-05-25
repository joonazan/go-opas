package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	summa := 0
	for i := 0; i <= n; i++ {
		summa += i * i
	}

	fmt.Println(summa)
}
