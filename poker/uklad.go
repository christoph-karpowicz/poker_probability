package poker

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
