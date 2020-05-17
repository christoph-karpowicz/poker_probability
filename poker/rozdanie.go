package poker

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
