package poker

import (
	"fmt"
	"sort"
)

type kartyGracza []*karta

func noweKartyGracza1Reka(kartaZReki *karta, kombinacjaKart []*karta) *kartyGracza {
	var karty5 kartyGracza = make(kartyGracza, 0)

	karty5 = append(karty5, kartaZReki)

	for _, krt := range kombinacjaKart {
		karty5 = append(karty5, krt)
	}

	// fmt.Println(wszystkieKarty)
	karty5.sortujWgFigury()
	// fmt.Println(wszystkieKarty)

	return &karty5
}

func noweKartyGraczaCalaReka(reka [2]*karta, kombinacjaKart []*karta) *kartyGracza {
	var karty5 kartyGracza = make(kartyGracza, 0)

	for _, krt := range reka {
		karty5 = append(karty5, krt)
	}

	for _, krt := range kombinacjaKart {
		karty5 = append(karty5, krt)
	}

	// fmt.Println(wszystkieKarty)
	karty5.sortujWgFigury()
	// fmt.Println(wszystkieKarty)

	return &karty5
}

func (kg *kartyGracza) sortujWgFigury() {
	sort.SliceStable(*kg, func(i, j int) bool {
		return znajdzIndexFigury((*kg)[i].figura) > znajdzIndexFigury((*kg)[j].figura)
	})
}

func (kg kartyGracza) String() string {
	result := ""

	for _, krt := range kg {
		result += fmt.Sprintf("%s ", krt)
	}

	return result
}

type gracz struct {
	reka [2]*karta
}

func (g *gracz) maPokeraKrolewskiego(karty *kartyGracza) bool {
	strit := g.maStrita(karty)
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

func (g *gracz) maPokera(karty *kartyGracza) bool {
	strit := g.maStrita(karty)
	if strit == nil {
		return false
	}

	if !strit.jestWJednymKolorze() {
		return false
	}

	return true
}

func (g *gracz) maKarete(karty *kartyGracza) bool {

	return true
}

func (g *gracz) maStrita(karty *kartyGracza) *uklad {
	var strit *uklad = &uklad{}
	var licznik int
	var indexOstatniejKarty int

	for i, krt := range *karty {
		var indexBiezacejKarty int = znajdzIndexFigury(krt.figura)
		// fmt.Println(indexOstatniejKarty)
		// fmt.Println(indexBiezacejKarty)

		if i > 0 {
			if indexBiezacejKarty == indexOstatniejKarty+1 {
				(*strit)[licznik] = krt
				(*strit)[licznik+1] = krt
				licznik++
			} else {
				licznik = 0
			}
		}

		indexOstatniejKarty = indexBiezacejKarty
	}

	if licznik >= 4 {
		return strit
	}
	return nil
}

func (g *gracz) sprawdzUklady(stol *stol, kombinacje3kart [][]*karta, kombinacje4kart [][]*karta) *uklad {
	ukl := &uklad{}

	ukladyGracza := make(map[string][]*uklad)
	for _, ukladKart := range ukladyKart {
		ukladyGracza[ukladKart] = make([]*uklad, 0)
	}

	for _, komb3 := range kombinacje3kart {

		kartyZReka := noweKartyGraczaCalaReka(g.reka, komb3)
		newUkl := uklad{}

		if g.maPokeraKrolewskiego(kartyZReka) {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["pokerKrolewski"] = append(ukladyGracza["pokerKrolewski"], &newUkl)
		} else if g.maPokera(kartyZReka) {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["poker"] = append(ukladyGracza["poker"], &newUkl)
		} else if g.maStrita(kartyZReka) != nil {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["strit"] = append(ukladyGracza["strit"], &newUkl)
		}
	}

	return ukl
}
