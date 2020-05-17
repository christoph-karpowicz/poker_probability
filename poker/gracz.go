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
	var strit *uklad = &uklad{}
	var licznik int
	var indexOstatniejKarty int

	for i, krt := range *kg {
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

type gracz struct {
	reka [2]*karta
}

func (g *gracz) sprawdzUklady(stol *stol, kombinacje3kart [][]*karta, kombinacje4kart [][]*karta) (string, *uklad) {
	var najwyzszyUkl *uklad
	var najwyzszyUklNazwa string
	znalezionoNajwyzszyUklad := false

	ukladyGracza := make(map[string][]*uklad)
	for _, ukladKart := range ukladyKart {
		ukladyGracza[ukladKart] = make([]*uklad, 0)
	}

	for _, komb3 := range kombinacje3kart {

		kartyZReka := noweKartyGraczaCalaReka(g.reka, komb3)
		newUkl := uklad{}

		if kartyZReka.maPokeraKrolewskiego() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["pokerKrolewski"] = append(ukladyGracza["pokerKrolewski"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "pokerKrolewski"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maPokera() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["poker"] = append(ukladyGracza["poker"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "poker"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maKarete() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["kareta"] = append(ukladyGracza["kareta"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "kareta"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maFula() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["ful"] = append(ukladyGracza["ful"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "ful"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maKolor() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["kolor"] = append(ukladyGracza["kolor"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "kolor"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maStrita() != nil {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["strit"] = append(ukladyGracza["strit"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "strit"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maTrojke() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["trojka"] = append(ukladyGracza["trojka"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "trojka"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maDwiePary() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["dwiePary"] = append(ukladyGracza["dwiePary"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "dwiePary"
				znalezionoNajwyzszyUklad = true
			}

		} else if kartyZReka.maPare() {
			newUkl = append(newUkl, (*kartyZReka)...)
			ukladyGracza["para"] = append(ukladyGracza["para"], &newUkl)

			if !znalezionoNajwyzszyUklad {
				najwyzszyUkl = &newUkl
				najwyzszyUklNazwa = "para"
				znalezionoNajwyzszyUklad = true
			}
		} else {
			najwyzszyUkl = &newUkl
			najwyzszyUklNazwa = "wysokaKarta"
		}
	}

	return najwyzszyUklNazwa, najwyzszyUkl
}
