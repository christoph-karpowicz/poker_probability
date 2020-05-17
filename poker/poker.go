package poker

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

func wyznaczKombinacjeKart(iloscKartWKombinacji int, karty [5]*karta) [][]*karta {
	result := make([][]*karta, 0)
	tmp := make([]*karta, iloscKartWKombinacji)

	kombinacjeKart(iloscKartWKombinacji, karty, 0, tmp, &result, 0)

	return result
}

func kombinacjeKart(iloscKartWKombinacji int, karty [5]*karta, index int, tmp []*karta, result *[][]*karta, i int) {
	if index == iloscKartWKombinacji {
		tmpCpy := make([]*karta, len(tmp))
		copy(tmpCpy, tmp)
		// fmt.Println(tmpCpy)
		*result = append(*result, tmpCpy)
		return
	}

	if i >= len(karty) {
		return
	}

	tmp[index] = karty[i]
	kombinacjeKart(iloscKartWKombinacji, karty, index+1, tmp, result, i+1)

	kombinacjeKart(iloscKartWKombinacji, karty, index, tmp, result, i+1)
}
