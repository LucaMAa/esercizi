package model

type Voto struct {
	Voto int
	CFU  int
}

type Registro map[int][]Voto
