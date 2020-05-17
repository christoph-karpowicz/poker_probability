package poker

import "fmt"

type stol struct {
	gracze                    []*gracz
	talia                     *talia
	licznikUkladowPrzedFlopem map[string]int
	licznikUkladowPoFlopie    map[string]int
	licznikUkladowPoTurnie    map[string]int
	licznikUkladowPoRiverze   map[string]int
}

func NowyStol(iloscGraczy int) *stol {
	nowaTalia := nowaTalia()
	gracze := make([]*gracz, 0)
	licznikUkladowPrzedFlopem := make(map[string]int)
	licznikUkladowPoFlopie := make(map[string]int)
	licznikUkladowPoTurnie := make(map[string]int)
	licznikUkladowPoRiverze := make(map[string]int)

	for i := 1; i <= iloscGraczy; i++ {
		gracze = append(gracze, &gracz{})
	}

	for _, ukladKart := range ukladyKart {
		licznikUkladowPrzedFlopem[ukladKart] = 0
		licznikUkladowPoFlopie[ukladKart] = 0
		licznikUkladowPoTurnie[ukladKart] = 0
		licznikUkladowPoRiverze[ukladKart] = 0
	}

	nowyStol := stol{
		talia:                     nowaTalia,
		gracze:                    gracze,
		licznikUkladowPrzedFlopem: licznikUkladowPrzedFlopem,
		licznikUkladowPoFlopie:    licznikUkladowPoFlopie,
		licznikUkladowPoTurnie:    licznikUkladowPoTurnie,
		licznikUkladowPoRiverze:   licznikUkladowPoRiverze,
	}

	return &nowyStol
}

func (s *stol) rozdaj() {
	s.talia.tasuj()
	s.talia.przeloz()

	rozdanie := noweRozdanie(s)
	rozdanie.rece()
	rozdanie.sprawdzUklady(s.licznikUkladowPrzedFlopem)

	rozdanie.flop()
	rozdanie.sprawdzUklady(s.licznikUkladowPoFlopie)

	rozdanie.turn()
	rozdanie.sprawdzUklady(s.licznikUkladowPoTurnie)

	rozdanie.river()
	rozdanie.sprawdzUklady(s.licznikUkladowPoRiverze)

	s.zbierzKarty(rozdanie)
}

func (s *stol) RozdajNrazy(iloscRozdan int) {
	for i := 0; i < iloscRozdan; i++ {
		s.rozdaj()
	}

	fmt.Println(s.licznikUkladowPrzedFlopem)
	fmt.Println(s.licznikUkladowPoFlopie)
	fmt.Println(s.licznikUkladowPoTurnie)
	fmt.Println(s.licznikUkladowPoRiverze)
}

func (s *stol) zbierzKarty(roz *rozdanie) {
	for _, krt := range roz.kartyWspolne {
		s.talia.odlozKarte(krt)
	}

	for _, gracz := range s.gracze {
		s.talia.odlozKarte(gracz.reka[0])
		s.talia.odlozKarte(gracz.reka[1])
	}
}
