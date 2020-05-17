package poker

import "fmt"

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

	for _, ukladKart := range ukladyKart {
		licznikUkladow[ukladKart] = 0
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
