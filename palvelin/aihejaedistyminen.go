package main

import "sort"

type AiheJaEdistyminen struct {
	Id string
	Aihe
	Tila uint8

	TäytetytVaatimukset, TäyttämättömätVaatimukset []string
}

type Aiheluettelo []AiheJaEdistyminen

func (a Aiheluettelo) Len() int {
	return len(a)
}

func (a Aiheluettelo) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var tilojenJärjestys = make([]uint8, TilojenMäärä)

func init() {
	tilojenJärjestys[Tehty] = 0
	tilojenJärjestys[Aloitettu] = 1
	tilojenJärjestys[EiTehty] = 2
}

func (a Aiheluettelo) Less(i, j int) bool {
	it := tilojenJärjestys[a[i].Tila]
	jt := tilojenJärjestys[a[j].Tila]

	return it < jt || it == jt && len(a[i].TäyttämättömätVaatimukset) < len(a[j].TäyttämättömätVaatimukset)
}

func (a Aiheluettelo) Lajittele() {
	sort.Sort(a)
}
