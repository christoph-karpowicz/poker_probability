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
	var kombinacje3kart [][]*karta
	kombinacje3kart = wyznaczKombinacjeKart(3, r.kartyWspolne)

	var kombinacje4kart [][]*karta
	kombinacje4kart = wyznaczKombinacjeKart(4, r.kartyWspolne)

	// for _, ko := range kombinacje4kart {
	// 	for _, k := range ko {
	// 		fmt.Print(k)
	// 	}
	// 	fmt.Println()
	// }

	for _, gracz := range r.stol.gracze {
		// wszystkieKarty := zlaczKartyGracza(gracz.reka, r.kartyWspolne)
		gracz.sprawdzUklady(r.stol, kombinacje3kart, kombinacje4kart)
	}
}
