package main

import "fmt"

type Vaihtoehto struct {
	kuvaus  string
	seuraus Tila
}

func valitse(vaihtoehdot []Vaihtoehto) Tila {

	for i, v := range vaihtoehdot {
		fmt.Printf("%c) %s\n", 'a'+i, v.kuvaus)
	}

	var sana string
	for len(sana) == 0 || sana[0] < 'a' || int(sana[0]-'a') >= len(vaihtoehdot) {
		fmt.Scan(&sana)
	}

	return vaihtoehdot[sana[0]-'a'].seuraus
}
