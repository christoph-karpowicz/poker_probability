package poker

import (
	"math/rand"
	"time"
)

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
