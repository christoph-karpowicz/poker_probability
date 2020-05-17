package poker

type gracz struct {
	reka [2]*karta
}

func (g *gracz) sprawdzUklady(kombinacje3kart [][]*karta, kombinacje4kart [][]*karta) string {
	var wszystkieUklady []string

	if kombinacje3kart != nil {
		for _, komb3 := range kombinacje3kart {
			kartyZReka := noweKartyGraczaCalaReka(g.reka, komb3)
			ukladNazwa := kartyZReka.szukajUkladow()

			wszystkieUklady = append(wszystkieUklady, ukladNazwa)
		}
	}

	if kombinacje4kart != nil {
		for _, komb4 := range kombinacje4kart {
			kartyZ1KartaReki := noweKartyGracza1Reka(g.reka[0], komb4)
			ukladNazwa := kartyZ1KartaReki.szukajUkladow()

			wszystkieUklady = append(wszystkieUklady, ukladNazwa)
		}
		for _, komb4 := range kombinacje4kart {
			kartyZ2KartaReki := noweKartyGracza1Reka(g.reka[1], komb4)
			ukladNazwa := kartyZ2KartaReki.szukajUkladow()

			wszystkieUklady = append(wszystkieUklady, ukladNazwa)
		}
	}

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
