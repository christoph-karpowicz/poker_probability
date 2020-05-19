package poker

import (
	"fmt"
	"sort"
)

type kartyGracza []*karta

// noweKartyGracza1Reka tworzy tablicę 5 kart
// użuwając 1 karty z reki gracza i danej kombinacji
// 4 kart wspólnych.
func noweKartyGracza1Reka(kartaZReki *karta, kombinacjaKart []*karta) *kartyGracza {
	var karty5 kartyGracza = make(kartyGracza, 0)

	karty5 = append(karty5, kartaZReki)

	for _, krt := range kombinacjaKart {
		karty5 = append(karty5, krt)
	}

	karty5.sortujWgFigury()

	return &karty5
}

// noweKartyGracza1Reka tworzy tablicę 5 kart
// użuwając reki gracza i danej kombinacji
// 3 kart wspólnych.
func noweKartyGraczaCalaReka(reka [2]*karta, kombinacjaKart []*karta) *kartyGracza {
	var karty5 kartyGracza = make(kartyGracza, 0)

	for _, krt := range reka {
		karty5 = append(karty5, krt)
	}

	for _, krt := range kombinacjaKart {
		karty5 = append(karty5, krt)
	}

	karty5.sortujWgFigury()

	return &karty5
}

// noweKartyGraczaTylkoReka tworzy "karty gracza"
// tylko z jego ręki.
func noweKartyGraczaTylkoReka(reka [2]*karta) *kartyGracza {
	var karty5 kartyGracza = make(kartyGracza, 0)

	for _, krt := range reka {
		karty5 = append(karty5, krt)
	}

	return &karty5
}

func (kg *kartyGracza) sortujWgFigury() {
	sort.SliceStable(*kg, func(i, j int) bool {
		return znajdzIndexFigury((*kg)[i].figura) < znajdzIndexFigury((*kg)[j].figura)
	})
}

func (kg kartyGracza) String() string {
	result := ""

	for _, krt := range kg {
		result += fmt.Sprintf("%s ", krt)
	}

	return result
}

// szukajUkladow szuka układów w 5 kartach.
func (kg *kartyGracza) szukajUkladow() string {
	if kg.maPokeraKrolewskiego() {
		return "pokerKrolewski"
	} else if kg.maPokera() {
		return "poker"
	} else if kg.maKarete() {
		return "kareta"
	} else if kg.maFula() {
		return "ful"
	} else if kg.maKolor() {
		return "kolor"
	} else if kg.maStrita() != nil {
		return "strit"
	} else if kg.maTrojke() {
		return "trojka"
	} else if kg.maDwiePary() {
		return "dwiePary"
	} else if kg.maPare() {
		return "para"
	} else {
		return "wysokaKarta"
	}
}

// szukajUkladowTylkoWRece szuka par w 2
// kartach.
func (kg *kartyGracza) szukajUkladowTylkoWRece() string {
	if kg.maPare() {
		return "para"
	} else {
		return "wysokaKarta"
	}
}

func (kg *kartyGracza) zliczWgFigury() map[string]int {
	figuryGracza := make(map[string]int)
	for _, figuraKart := range figuryKart {
		figuryGracza[figuraKart] = 0
	}

	for _, krt := range *kg {
		figuryGracza[krt.figura]++
	}

	return figuryGracza
}

func (kg *kartyGracza) zliczWgKoloru() map[string]int {
	koloryGracza := make(map[string]int)
	for _, kolorKart := range koloryKart {
		koloryGracza[kolorKart] = 0
	}

	for _, krt := range *kg {
		koloryGracza[krt.kolor]++
	}

	return koloryGracza
}

func (kg *kartyGracza) maPokeraKrolewskiego() bool {
	strit := kg.maStrita()
	if strit == nil {
		return false
	}

	if !strit.jestWJednymKolorze() {
		return false
	}

	if (*strit)[0].figura != "10" {
		return false
	}

	return true
}

func (kg *kartyGracza) maPokera() bool {
	strit := kg.maStrita()
	if strit == nil {
		return false
	}

	if !strit.jestWJednymKolorze() {
		return false
	}

	return true
}

func (kg *kartyGracza) maKarete() bool {
	figury := kg.zliczWgFigury()

	for _, iloscFigur := range figury {
		if iloscFigur == 4 {
			return true
		}
	}

	return false
}

func (kg *kartyGracza) maFula() bool {
	figury := kg.zliczWgFigury()
	trzyTakieSame := ""

	for figura, iloscFigur := range figury {
		if iloscFigur == 3 {
			trzyTakieSame = figura
		}
	}

	if trzyTakieSame == "" {
		return false
	}

	for figura, iloscFigur := range figury {
		if figura != trzyTakieSame {
			if iloscFigur >= 2 {
				return true
			}
		}
	}

	return false
}

func (kg *kartyGracza) maKolor() bool {
	kolory := kg.zliczWgKoloru()

	for _, iloscKolorow := range kolory {
		if iloscKolorow == 5 {
			return true
		}
	}

	return false
}

func (kg *kartyGracza) maStrita() *uklad {
	var strit uklad = make(uklad, 5)
	var licznik int
	var indexOstatniejKarty int

	for i, krt := range *kg {
		var indexBiezacejKarty int = znajdzIndexFigury(krt.figura)

		if i > 0 {
			if indexBiezacejKarty == indexOstatniejKarty+1 {
				strit[licznik] = krt
				strit[licznik+1] = krt
				licznik++
			} else {
				return nil
			}
		}

		indexOstatniejKarty = indexBiezacejKarty
	}

	return &strit
}

func (kg *kartyGracza) maTrojke() bool {
	figury := kg.zliczWgFigury()

	for _, iloscFigur := range figury {
		if iloscFigur == 3 {
			return true
		}
	}

	return false
}

func (kg *kartyGracza) maDwiePary() bool {
	figury := kg.zliczWgFigury()
	dwieTakieSame := 0

	for _, iloscFigur := range figury {
		if iloscFigur == 2 {
			dwieTakieSame++
		}
	}

	if dwieTakieSame >= 2 {
		return true
	}
	return false
}

func (kg *kartyGracza) maPare() bool {
	figury := kg.zliczWgFigury()

	for _, iloscFigur := range figury {
		if iloscFigur == 2 {
			return true
		}
	}

	return false
}
