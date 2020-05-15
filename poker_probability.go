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
	reka []*karta
}

type stol struct {
	gracze []gracz
	talia  talia
}

type talia struct {
	karty []karta
}

func nowaTalia() *talia {
	nowaTalia := talia{
		karty: make([]karta, 0),
	}

	for _, kolor := range koloryKart {
		for _, figura := range figuryKart {
			nowaTalia.karty = append(nowaTalia.karty, karta{
				kolor:  kolor,
				figura: figura,
			})
		}
	}

	return &nowaTalia
}

func (t *talia) tasuj() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(t.karty), func(i, j int) { t.karty[i], t.karty[j] = t.karty[j], t.karty[i] })
}

func (t *talia) przeloz() {
	iloscKartWCzesci := rand.Intn(30-20) + 20

	czesc1 := t.karty[0:iloscKartWCzesci]
	czesc2 := t.karty[iloscKartWCzesci:len(t.karty)]

	t.karty = append(czesc2, czesc1...)
}

func (t *talia) rozdaj(iloscGraczy int) {
}

type karta struct {
	kolor  string
	figura string
}

func main() {
	fmt.Println("Generowanie losowych rozdań w pokera i empiryczne wyznaczenie prawdopodobieństwa wszystkich konfiguracji.")

	talia := nowaTalia()
	talia.tasuj()
	talia.przeloz()
	fmt.Println(talia.karty)
}
