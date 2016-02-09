package main

import "testing"

func TestRecycler(t *testing.T) {
	var a []int
	var b [][2]float32

	k := UusiKierrättäjä(&a, &b)

	const num = 17

	idt := make(map[int]bool)
	for i := 0; i < 17; i++ {
		idt[k.Varaa()] = true
	}
	if len(idt) != num {
		t.Fatal("Annetut paikat eivät olleet uniikkeja.")
	}

	if len(a) != num || len(b) != num {
		t.Fatal("Slicet eivät ole oikean kokoisia.")
	}
}
