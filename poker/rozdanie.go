package poker

type rozdanie struct {
	stol         *stol
	kartyWspolne []*karta
}

func noweRozdanie(stol *stol) *rozdanie {
	noweRozdanie := rozdanie{
		stol:         stol,
		kartyWspolne: make([]*karta, 0),
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
		r.kartyWspolne = append(r.kartyWspolne, r.stol.talia.pobierzOstatniaKarte())
	}
}

func (r *rozdanie) turn() {
	r.kartyWspolne = append(r.kartyWspolne, r.stol.talia.pobierzOstatniaKarte())
}

func (r *rozdanie) river() {
	r.kartyWspolne = append(r.kartyWspolne, r.stol.talia.pobierzOstatniaKarte())
}

func (r *rozdanie) oddajKarty() {

}

func (r *rozdanie) sprawdzUklady(licznik map[string]int) {
	iloscKartWspolnych := len(r.kartyWspolne)

	var kombinacje3kart [][]*karta
	if iloscKartWspolnych >= 4 {
		kombinacje3kart = wyznaczKombinacjeKart(3, r.kartyWspolne)
	} else if iloscKartWspolnych == 3 {
		kombinacje3kart = make([][]*karta, 1)
		kombinacje3kart[0] = r.kartyWspolne
	} else {
		kombinacje3kart = nil
	}

	var kombinacje4kart [][]*karta
	if iloscKartWspolnych == 5 {
		kombinacje4kart = wyznaczKombinacjeKart(4, r.kartyWspolne)
	} else if iloscKartWspolnych == 4 {
		kombinacje4kart = make([][]*karta, 1)
		kombinacje4kart[0] = r.kartyWspolne
	} else {
		kombinacje4kart = nil
	}

	// for _, ko := range kombinacje4kart {
	// 	for _, k := range ko {
	// 		fmt.Print(k)
	// 	}
	// 	fmt.Println()
	// }

	for _, gracz := range r.stol.gracze {
		najwyzszyUkladNazwa := gracz.sprawdzUklady(kombinacje3kart, kombinacje4kart)
		licznik[najwyzszyUkladNazwa]++
	}
}
