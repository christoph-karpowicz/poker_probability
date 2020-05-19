package poker

const iloscKartWTalii = 52

var koloryKart []string = []string{"pik", "kier", "trefl", "karo"}
var figuryKart []string = []string{"as", "król", "dama", "walet", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
var ukladyKart []string = []string{"pokerKrolewski", "poker", "kareta", "ful", "kolor", "strit", "trojka", "dwiePary", "para", "wysokaKarta"}
var ukladyKartPelneNazwy []string = []string{"poker królewski", "poker", "kareta", "ful", "kolor", "strit", "trójka", "dwie pary", "para", "wysoka karta"}

func znajdzIndexFigury(figura string) int {
	for i, fgra := range figuryKart {
		if figura == fgra {
			return i
		}
	}

	return 0
}

func znajdzIndexUkladu(uklad string) int {
	for i, ukl := range ukladyKart {
		if uklad == ukl {
			return i
		}
	}

	return 0
}

// wyznaczKombinacjeKart wyznacza wszystkie możliwe
// kombinacje podanej ilości kart wspólnych znajdujących
// się na stole.
func wyznaczKombinacjeKart(iloscKartWKombinacji int, karty []*karta) [][]*karta {
	wszystkieKombinacje := make([][]*karta, 0)
	tmp := make([]*karta, iloscKartWKombinacji)

	kombinacjeKart(iloscKartWKombinacji, karty, 0, tmp, &wszystkieKombinacje, 0)

	return wszystkieKombinacje
}

// kombinacjeKart rekurencyjnie dodaje kolejne
// kombinacje zapisywane w tablicy tymczasowej
// do przekazanej jej, tablicy wszystkich kombinacji.
func kombinacjeKart(iloscKartWKombinacji int, karty []*karta, index int, tmp []*karta, wszystkieKombinacje *[][]*karta, i int) {
	if index == iloscKartWKombinacji {
		tmpCpy := make([]*karta, len(tmp))

		copy(tmpCpy, tmp)
		*wszystkieKombinacje = append(*wszystkieKombinacje, tmpCpy)
		return
	}

	if i >= len(karty) {
		return
	}

	tmp[index] = karty[i]
	kombinacjeKart(iloscKartWKombinacji, karty, index+1, tmp, wszystkieKombinacje, i+1)

	kombinacjeKart(iloscKartWKombinacji, karty, index, tmp, wszystkieKombinacje, i+1)
}
