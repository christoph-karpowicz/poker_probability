package poker

import (
	"fmt"
	"sort"
)

type kartyGracza []*karta

func zlaczKartyGracza(reka [2]*karta, kartyWspolne [5]*karta) *kartyGracza {
	var wszystkieKarty kartyGracza = make(kartyGracza, 0)
	for _, krt := range kartyWspolne {
		wszystkieKarty = append(wszystkieKarty, krt)
	}
	for _, krt := range reka {
		wszystkieKarty = append(wszystkieKarty, krt)
	}

	// fmt.Println(wszystkieKarty)
	wszystkieKarty.sortujWgFigury()
	// fmt.Println(wszystkieKarty)

	return &wszystkieKarty
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

func (g *gracz) maPokeraKrolewskiego(wszystkieKarty *kartyGracza) bool {
	strit := g.maStrita(wszystkieKarty)
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

func (g *gracz) maPokera(wszystkieKarty *kartyGracza) bool {
	strit := g.maStrita(wszystkieKarty)
	if strit == nil {
		return false
	}

	if !strit.jestWJednymKolorze() {
		return false
	}

	return true
}

func (g *gracz) maKarete(wszystkieKarty *kartyGracza) bool {

	return true
}

func (g *gracz) maStrita(wszystkieKarty *kartyGracza) *uklad {
	var strit *uklad = &uklad{}
	var licznik int
	var indexOstatniejKarty int

	for i, krt := range *wszystkieKarty {
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
