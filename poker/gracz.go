package poker

type gracz struct {
	reka [2]*karta
}

func (g *gracz) sprawdzKombinacje(kombinacje [][]*karta) []string {
	var wszystkieUklady []string

	for _, komb := range kombinacje {
		if len(komb) == 3 {
			kartyZReka := noweKartyGraczaCalaReka(g.reka, komb)
			wszystkieUklady = append(wszystkieUklady, kartyZReka.szukajUkladow())
		} else {
			kartyZ1KartaReki := noweKartyGracza1Reka(g.reka[0], komb)
			wszystkieUklady = append(wszystkieUklady, kartyZ1KartaReki.szukajUkladow())

			kartyZ2KartaReki := noweKartyGracza1Reka(g.reka[1], komb)
			wszystkieUklady = append(wszystkieUklady, kartyZ2KartaReki.szukajUkladow())
		}
	}

	return wszystkieUklady
}

func (g *gracz) sprawdzUklady(kombinacje3kart [][]*karta, kombinacje4kart [][]*karta) string {
	var wszystkieUklady []string
	var wszystkieKombinacje [][]*karta = kombinacje3kart

	wszystkieKombinacje = append(wszystkieKombinacje, kombinacje4kart...)

	wszystkieUklady = g.sprawdzKombinacje(wszystkieKombinacje)

	if kombinacje3kart == nil && kombinacje4kart == nil {
		kartyTylkoReka := noweKartyGraczaTylkoReka(g.reka)
		ukladNazwa := kartyTylkoReka.szukajUkladowTylkoWRece()

		wszystkieUklady = append(wszystkieUklady, ukladNazwa)
	}

	var indexNajwyzszegoUkladu int

	for i, ukl := range wszystkieUklady {
		if i == 0 {
			indexNajwyzszegoUkladu = znajdzIndexUkladu(ukl)
		} else {
			if indexBizacegoUkladu := znajdzIndexUkladu(ukl); indexNajwyzszegoUkladu > indexBizacegoUkladu {
				indexNajwyzszegoUkladu = indexBizacegoUkladu
			}
		}

	}

	return ukladyKart[indexNajwyzszegoUkladu]
}
