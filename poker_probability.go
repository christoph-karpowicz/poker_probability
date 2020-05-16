package main

import (
	"fmt"
	"math/rand"
	"time"
)

const iloscKartWTalii = 52

var koloryKart []string = []string{"pik", "kier", "trefl", "karo"}
var figuryKart []string = []string{"as", "król", "dama", "walet", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

type gracz struct {
	reka [2]*karta
}

type stol struct {
	gracze []gracz
	talia  *talia
}

func nowyStol(iloscGraczy int) *stol {
	nowaTalia := nowaTalia()
	gracze := make([]gracz, 0)

	for i := 1; i <= iloscGraczy; i++ {
		gracze = append(gracze, gracz{})
	}

	nowyStol := stol{
		talia:  nowaTalia,
		gracze: gracze,
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
			nowaTalia.karty = append(nowaTalia.karty, &karta{
				kolor:  kolor,
				figura: figura,
			})
		}
	}

	return &nowaTalia
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
	noweRozdanie := rozdanie{
		stol: stol,
	}

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

func main() {
	fmt.Println("Generowanie losowych rozdań w pokera i empiryczne wyznaczenie prawdopodobieństwa wszystkich konfiguracji.")

	stol10 := nowyStol(10)
	fmt.Println(len(stol10.talia.karty))
	stol10.rozdajNrazy(1)
	fmt.Println(len(stol10.talia.karty))
}
