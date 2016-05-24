package main

import (
	"github.com/joonazan/vec2"
	"sort"
)

var törmääjät = TeeTörmääjät()

type Törmääjät struct {
	Komponentti
	Säde   []float64
	Tyyppi []uint
}

func TeeTörmääjät() (t Törmääjät) {
	t.Komponentti = maailma.LuoKomponentti(&kuolleetIDt, &t.Säde, &t.Tyyppi)
	return
}

func (t Törmääjät) Paikka(i int) vec2.Vector {
	return paikat[törmääjälle.Vanhemmat[j]]
}

const piirräTörmäysympyrät = true

type päätepiste struct {
	// id:n ensimmäinen bitti kertoo onko kyseessä alkupiste vai loppupiste
	id     uint32
	Paikka float64
}

func (p päätepiste) Id() int {
	return p.id &^ (1 << 31)
}

func (p päätepiste) OnAlkupiste() bool {
	return (p.id >> 31) == 0
}

type päätepisteet []päätepiste

func (t päätepisteet) Less(i, j int) bool {
	return t[i].Paikka.x < t[j].Paikka.x
}

func (t päätepisteet) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t päätepisteet) Len() int {
	return len(t)
}

func (t Törmääjät) LöydäTörmäykset() {
	pisteet := make(päätepisteet, 2*len(t.Säde))

	for i := range t.Säde {
		keskipiste := t.Paikka(i)
		pisteet = append(pisteet,
			päätepiste{i, keskipiste - t.Säde[i]},
			päätepiste{i | (1 << 31), keskipiste + t.Säde[i]},
		)
	}

	sort.Sort(pisteet)

	var aktiiviset []int
	for _, piste := range pisteet {
		if piste.OnAlkupiste() {
			for _, i := range aktiiviset {
				minimietäisyys := t.Säde[i] + t.Säde[piste.Id()]
				paikkojenEro := t.Paikka(i).Minus(t.Paikka(piste.Id))
				if paikkojenEro.LengthSquared() < minimietäisyys*minimietäisyys {
					törmäykset = append(törmäykset, [2]int{piste.Id(), i})
				}
			}
			aktiiviset = append(aktiiviset, piste.Id())
		} else {
			for i, id := range aktiiviset {
				if id == piste.Id() {
					viimeinen := len(aktiiviset) - 1
					aktiiviset[i] = aktiiviset[viimeinen]
					aktiiviset = aktiiviset[:viimeinen]
					break
				}
			}
		}
	}
}
