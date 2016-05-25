package main

import "fmt"

func main() {
	var input, input2 string
	fmt.Scan(&input, &input2)

	if input != "moi" && input2 == "apua" {
		fmt.Println("kuole")
	}
}
