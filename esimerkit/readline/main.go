package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	lukija := bufio.NewReader(os.Stdin)

	fmt.Print("Kirjoita jotain tekstiä! ")
	rivi, _ := lukija.ReadString('\n')

	fmt.Println("kirjoitit:", rivi)
}
