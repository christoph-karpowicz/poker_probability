package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

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

type uklad []*karta

func (u *uklad) jestWJednymKolorze() bool {
	var pierwszyKolor string

	for i, krt := range *u {
		if i == 0 {
			pierwszyKolor = krt.kolor
			continue
		}

		if krt.kolor != pierwszyKolor {
			return false
		}
	}

	return true
}

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

type stol struct {
	gracze         []*gracz
	talia          *talia
	licznikUkladow map[string]int
}

func nowyStol(iloscGraczy int) *stol {
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

func (s *stol) rozdajNrazy(iloscRozdan int) {
	for i := 0; i < iloscRozdan; i++ {
		s.rozdaj()
	}
}

type talia struct {
	karty []*karta
}

func nowaTalia() *talia {
	nowaTalia := talia{
		karty: make([]*karta, 0),
	}

	for _, kolor := range koloryKart {
		for _, figura := range figuryKart {
			nowaTalia.karty = append(nowaTalia.karty, &karta{kolor, figura})
		}
	}

	return &nowaTalia
}

func (t *talia) odlozKarte(krt *karta) {
	t.karty = append(t.karty, krt)
}

func (t *talia) pobierzOstatniaKarte() *karta {
	karta := t.karty[len(t.karty)-1]
	t.karty = t.karty[:len(t.karty)-1]

	return karta
}

func (t *talia) przeloz() {
	iloscKartWCzesci := rand.Intn(30-20) + 20

	czesc1 := t.karty[0:iloscKartWCzesci]
	czesc2 := t.karty[iloscKartWCzesci:len(t.karty)]

	t.karty = append(czesc2, czesc1...)
}

func (t *talia) tasuj() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(t.karty), func(i, j int) { t.karty[i], t.karty[j] = t.karty[j], t.karty[i] })
}

type karta struct {
	kolor  string
	figura string
}

type rozdanie struct {
	stol         *stol
	kartyWspolne [5]*karta
}

func noweRozdanie(stol *stol) *rozdanie {
	noweRozdanie := rozdanie{stol: stol}

	return &noweRozdanie
}

func (r *rozdanie) rece() {
	for _, gracz := range r.stol.gracze {
		gracz.reka[0] = r.stol.talia.pobierzOstatniaKarte()
		gracz.reka[1] = r.stol.talia.pobierzOstatniaKarte()
	}
}

func (r *rozdanie) flop() {
	for i := 0; i < 3; i++ {
		r.kartyWspolne[i] = r.stol.talia.pobierzOstatniaKarte()
	}
}

func (r *rozdanie) turn() {
	r.kartyWspolne[3] = r.stol.talia.pobierzOstatniaKarte()
}

func (r *rozdanie) river() {
	r.kartyWspolne[4] = r.stol.talia.pobierzOstatniaKarte()
}

func (r *rozdanie) oddajKarty() {

}

func (r *rozdanie) sprawdzUklady() {
	for _, gracz := range r.stol.gracze {
		wszystkieKarty := zlaczKartyGracza(gracz.reka, r.kartyWspolne)

		if gracz.maPokeraKrolewskiego(wszystkieKarty) {
			r.stol.licznikUkladow["pokerKrolewski"]++
		} else if gracz.maPokera(wszystkieKarty) {
			r.stol.licznikUkladow["poker"]++
		} else if gracz.maStrita(wszystkieKarty) != nil {
			r.stol.licznikUkladow["strit"]++
		}
	}
}

func main() {
	fmt.Println("Generowanie losowych rozdań w pokera i empiryczne wyznaczenie prawdopodobieństwa wszystkich konfiguracji.")

	stol10 := nowyStol(10)
	stol10.rozdajNrazy(1)
	fmt.Println(stol10.licznikUkladow)
}
