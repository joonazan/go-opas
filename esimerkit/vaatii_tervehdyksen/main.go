package main

import "fmt"

func main() {
	for kirjoitettu() != "hei" {
		fmt.Println("Et tervehtinyt!")
	}
	fmt.Println("Hei hei!")
}

func kirjoitettu() string {
	var s string
	fmt.Scan(&s)
	return s
}
