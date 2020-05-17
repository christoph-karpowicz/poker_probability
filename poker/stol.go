package poker

import "fmt"

const iloscKartWTalii = 52

var koloryKart []string = []string{"pik", "kier", "trefl", "karo"}
var figuryKart []string = []string{"as", "król", "dama", "walet", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
var ukladyKart []string = []string{"pokerKrolewski", "poker", "kareta", "ful", "kolor", "strit", "trojka", "dwiePary", "para", "wysokaKarta"}

func znajdzIndexFigury(figura string) int {
	for i, fgra := range figuryKart {
		if figura == fgra {
			return i
		}
	}

	return 0
}

type stol struct {
	gracze         []*gracz
	talia          *talia
	licznikUkladow map[string]int
}

func NowyStol(iloscGraczy int) *stol {
	nowaTalia := nowaTalia()
	gracze := make([]*gracz, 0)
	licznikUkladow := make(map[string]int)

	for i := 1; i <= iloscGraczy; i++ {
		gracze = append(gracze, &gracz{})
	}

	for _, uklad := range ukladyKart {
		licznikUkladow[uklad] = 0
	}

	nowyStol := stol{
		talia:          nowaTalia,
		gracze:         gracze,
		licznikUkladow: licznikUkladow,
	}

	return &nowyStol
}

func (s *stol) rozdaj() {
	s.talia.tasuj()
	s.talia.przeloz()

	rozdanie := noweRozdanie(s)
	rozdanie.rece()
	rozdanie.flop()
	rozdanie.turn()
	rozdanie.river()

	rozdanie.sprawdzUklady()
}

func (s *stol) RozdajNrazy(iloscRozdan int) {
	for i := 0; i < iloscRozdan; i++ {
		s.rozdaj()
	}

	fmt.Println(s.licznikUkladow)
}
