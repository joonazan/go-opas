package main

type Aihe struct {
	Nimi string
}

type AiheJaEdistyminen struct {
	Id string
	Aihe
	Tila uint8
}
