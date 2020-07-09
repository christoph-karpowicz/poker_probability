package main

import (
	"fmt"
	"sort"
)

type playersCards []*card

// newPlayersCards1Hand creates a slice with 5 cards
// using 1 card from the hand and 4 from the given
// card combination.
func newPlayersCards1Hand(kartaZReki *card, cardCombination []*card) *playersCards {
	var cards playersCards = make(playersCards, 0)

	cards = append(cards, kartaZReki)

	for _, crd := range cardCombination {
		cards = append(cards, crd)
	}

	cards.sortByValue()

	return &cards
}

// newPlayersCardsWholeHand creates a slice with 5 cards
// using the player's hand and 3 cards from the given
// card combination.
func newPlayersCardsWholeHand(hand [2]*card, cardCombination []*card) *playersCards {
	var cards playersCards = make(playersCards, 0)

	for _, crd := range hand {
		cards = append(cards, crd)
	}

	for _, crd := range cardCombination {
		cards = append(cards, crd)
	}

	cards.sortByValue()

	return &cards
}

// newPlayersCardsOnlyFlop creates playersCards
// only from the player's flop.
func newPlayersCardsOnlyFlop(hand [2]*card) *playersCards {
	var cards playersCards = make(playersCards, 0)

	for _, crd := range hand {
		cards = append(cards, crd)
	}

	return &cards
}

func (pc *playersCards) sortByValue() {
	sort.SliceStable(*pc, func(i, j int) bool {
		return findCardValuesIndex((*pc)[i].value) < findCardValuesIndex((*pc)[j].value)
	})
}

func (pc playersCards) String() string {
	result := ""

	for _, crd := range pc {
		result += fmt.Sprintf("%s ", crd)
	}

	return result
}

// findHands looks for hand in 5 cards.
func (pc *playersCards) findHands() string {
	if pc.hasRoyalFlush() {
		return "royalFlush"
	} else if pc.hasStraightFlush() {
		return "straightFlush"
	} else if pc.hasFourOfAKind() {
		return "fourOfAkind"
	} else if pc.hasFullHouse() {
		return "fullHouse"
	} else if pc.hasFlush() {
		return "flush"
	} else if pc.hasStraight() != nil {
		return "straight"
	} else if pc.hasThreeOfAKind() {
		return "threeOfAkind"
	} else if pc.hasTwoPairs() {
		return "twoPairs"
	} else if pc.hasPair() {
		return "pair"
	} else {
		return "highCard"
	}
}

// findHandsOnlyFlop checks if 2 cards
// are a pair.
func (pc *playersCards) findHandsOnlyFlop() string {
	if pc.hasPair() {
		return "pair"
	} else {
		return "highCard"
	}
}

func (pc *playersCards) countByValue() map[string]int {
	playersValues := make(map[string]int)
	for _, cardValue := range cardValues {
		playersValues[cardValue] = 0
	}

	for _, crd := range *pc {
		playersValues[crd.value]++
	}

	return playersValues
}

func (pc *playersCards) countBySuit() map[string]int {
	playersSuits := make(map[string]int)
	for _, kolorKart := range cardSuits {
		playersSuits[kolorKart] = 0
	}

	for _, crd := range *pc {
		playersSuits[crd.suit]++
	}

	return playersSuits
}

func (pc *playersCards) hasRoyalFlush() bool {
	straight := pc.hasStraight()
	if straight == nil {
		return false
	}

	if !straight.areOfSameSuit() {
		return false
	}

	if (*straight)[0].value != "10" {
		return false
	}

	return true
}

func (pc *playersCards) hasStraightFlush() bool {
	straight := pc.hasStraight()
	if straight == nil {
		return false
	}

	if !straight.areOfSameSuit() {
		return false
	}

	return true
}

func (pc *playersCards) hasFourOfAKind() bool {
	cardValues := pc.countByValue()

	for _, nValues := range cardValues {
		if nValues == 4 {
			return true
		}
	}

	return false
}

func (pc *playersCards) hasFullHouse() bool {
	cardValues := pc.countByValue()
	threeOfSameValue := ""

	for cardValue, nValues := range cardValues {
		if nValues == 3 {
			threeOfSameValue = cardValue
		}
	}

	if threeOfSameValue == "" {
		return false
	}

	for cardValue, nValues := range cardValues {
		if cardValue != threeOfSameValue {
			if nValues >= 2 {
				return true
			}
		}
	}

	return false
}

func (pc *playersCards) hasFlush() bool {
	kolory := pc.countBySuit()

	for _, iloscKolorow := range kolory {
		if iloscKolorow == 5 {
			return true
		}
	}

	return false
}

func (pc *playersCards) hasStraight() *hand {
	var straight hand = make(hand, 5)
	var counter int
	var lastCardsIndex int

	for i, crd := range *pc {
		var currentCardsIndex int = findCardValuesIndex(crd.value)

		if i > 0 {
			if currentCardsIndex == lastCardsIndex+1 {
				straight[counter] = crd
				straight[counter+1] = crd
				counter++
			} else {
				return nil
			}
		}

		lastCardsIndex = currentCardsIndex
	}

	return &straight
}

func (pc *playersCards) hasThreeOfAKind() bool {
	cardValues := pc.countByValue()

	for _, nValues := range cardValues {
		if nValues == 3 {
			return true
		}
	}

	return false
}

func (pc *playersCards) hasTwoPairs() bool {
	cardValues := pc.countByValue()
	twoOfSameValue := 0

	for _, nValues := range cardValues {
		if nValues == 2 {
			twoOfSameValue++
		}
	}

	if twoOfSameValue >= 2 {
		return true
	}
	return false
}

func (pc *playersCards) hasPair() bool {
	cardValues := pc.countByValue()

	for _, nValues := range cardValues {
		if nValues == 2 {
			return true
		}
	}

	return false
}
