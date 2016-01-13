package main

type AiheJaEdistyminen struct {
	Id string
	Aihe
	Tila uint8
}

func lajittele(a []AiheJaEdistyminen) []AiheJaEdistyminen {

	tulos := make([]AiheJaEdistyminen, 0, len(a))

	for _, tila := range []uint8{Tehty, Aloitettu, EiTehty} {
		for _, aihe := range a {
			if aihe.Tila == tila {
				tulos = append(tulos, aihe)
			}
		}
	}

	return tulos
}
