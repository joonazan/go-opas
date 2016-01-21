package main

import "fmt"

func main() {
	ruuhka()
	fmt.Println("IIK!")
	ruuhka()
	melua()
}

func ruuhka() {
	melua()
	melua()
	melua()
}

func melua() {
	fmt.Println("TÖÖT!")
}
