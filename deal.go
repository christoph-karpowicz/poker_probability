package main

import (
	"sync"
)

type deal struct {
	table          *table
	communityCards []*card
}

func newDeal(table *table) *deal {
	newDeal := deal{
		table:          table,
		communityCards: make([]*card, 0),
	}

	return &newDeal
}

// dealFlops deals 2 cards for each player.
func (r *deal) dealFlops() {
	for _, player := range r.table.players {
		player.flop[0] = r.table.deck.drawLastCard()
		player.flop[1] = r.table.deck.drawLastCard()

		r.table.handCounter++
	}
}

func (r *deal) flop() {
	for i := 0; i < 3; i++ {
		r.communityCards = append(r.communityCards, r.table.deck.drawLastCard())
	}
}

func (r *deal) turn() {
	r.communityCards = append(r.communityCards, r.table.deck.drawLastCard())
}

func (r *deal) river() {
	r.communityCards = append(r.communityCards, r.table.deck.drawLastCard())
}

// findHands looks for hands in all combinations of
// players hands and all community cards.
func (r *deal) findHands(counter map[string]int) {
	numberOfCommunityCards := len(r.communityCards)

	var combinationsOf3Cards [][]*card
	if numberOfCommunityCards >= 4 {
		combinationsOf3Cards = get3CardCombination(3, r.communityCards)
	} else if numberOfCommunityCards == 3 {
		combinationsOf3Cards = make([][]*card, 1)
		combinationsOf3Cards[0] = r.communityCards
	} else {
		combinationsOf3Cards = nil
	}

	var combinationsOf4Cards [][]*card
	if numberOfCommunityCards == 5 {
		combinationsOf4Cards = get3CardCombination(4, r.communityCards)
	} else if numberOfCommunityCards == 4 {
		combinationsOf4Cards = make([][]*card, 1)
		combinationsOf4Cards[0] = r.communityCards
	} else {
		combinationsOf4Cards = nil
	}

	var wg sync.WaitGroup
	var lock = sync.RWMutex{}

	for _, player := range r.table.players {

		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			bestHandName := player.findHands(combinationsOf3Cards, combinationsOf4Cards)

			lock.Lock()
			counter[bestHandName]++
			lock.Unlock()

			wg.Done()
		}(&wg)

	}

	wg.Wait()
}
