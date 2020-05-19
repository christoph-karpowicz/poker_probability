package poker

import (
	"fmt"
	"strconv"
	"strings"
)

type stol struct {
	gracze                    []*gracz
	talia                     *talia
	licznikRak                int
	licznikRozdan             int
	licznikUkladowPrzedFlopem map[string]int
	licznikUkladowPoFlopie    map[string]int
	licznikUkladowPoTurnie    map[string]int
	licznikUkladowPoRiverze   map[string]int
}

// NowyStol tworzy nowy stół z daną ilością
// graczy
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

// ObliczPrawdopodobienstwa oblicza i wyświetla prawdopodobieństwa
// układów po N rozdań na podstawie liczników z poszczególych
// etapów gry.
func (s *stol) ObliczPrawdopodobienstwa() {
	s.obliczPrawdopodobienstwaNaEtapie("przed flopem", 2, s.licznikUkladowPrzedFlopem)
	s.obliczPrawdopodobienstwaNaEtapie("po flopie", 5, s.licznikUkladowPoFlopie)
	s.obliczPrawdopodobienstwaNaEtapie("po turnie", 6, s.licznikUkladowPoTurnie)
	s.obliczPrawdopodobienstwaNaEtapie("po riverze", 7, s.licznikUkladowPoRiverze)
}

func (s *stol) obliczPrawdopodobienstwaNaEtapie(nazwaEtapu string, iloscKart int, licznik map[string]int) {
	fmt.Println()
	fmt.Println("Prawdopodobieństwa wystąpień układów kart " + nazwaEtapu + ":")
	fmt.Println(strings.Repeat("_", 66))
	fmt.Printf("%-30s| %-34s|\n", "Nazwa układu", "Prawdopodobieństwo wystąpienia (%)")
	fmt.Println(strings.Repeat("-", 67))

	for i, ukl := range ukladyKart {
		if iloscKart == 2 && i != 8 && i != 9 {
			continue
		}

		iloscWystapien := licznik[ukl]
		var prawdopodobienstwo float64 = (float64(iloscWystapien) / float64(s.licznikRak)) * float64(100)

		fmt.Printf("%-30s| %-34.3f|\n", ukladyKartPelneNazwy[i], prawdopodobienstwo)
	}
	fmt.Println(strings.Repeat("-", 67))

	fmt.Printf("%-42s: %v\n", "Ilość kart dostępna dla jednego gracza", strconv.FormatInt(int64(iloscKart), 10))
	fmt.Printf("%-42s: %v\n", "Ilość graczy", strconv.FormatInt(int64(len(s.gracze)), 10))
	fmt.Printf("%-42s: %v\n", "Ilość rozdań", strconv.FormatInt(int64(s.licznikRozdan), 10))
	fmt.Printf("%-42s: %v\n", "Ilość rąk", strconv.FormatInt(int64(s.licznikRak), 10))
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

	s.licznikRozdan++
}

func (s *stol) RozdajNrazy(iloscRozdan int) {
	for i := 0; i < iloscRozdan; i++ {
		s.rozdaj()
	}
}

// zbierzKarty pobiera rozdane karty
// spowrotem do talii.
func (s *stol) zbierzKarty(roz *rozdanie) {
	for _, krt := range roz.kartyWspolne {
		s.talia.odlozKarte(krt)
	}

	for _, gracz := range s.gracze {
		s.talia.odlozKarte(gracz.reka[0])
		s.talia.odlozKarte(gracz.reka[1])
	}
}
